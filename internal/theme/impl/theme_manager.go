package impl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/utils/theme"
	"github.com/Done-0/jank/pkg/theme/consts"
)

// ThemeManagerImpl 主题管理器实现
type ThemeManagerImpl struct {
	themes      map[string]*ThemeInfo // 主题信息映射
	activeTheme string                // 当前激活的主题ID
	mu          sync.RWMutex          // 并发安全锁
}

// NewThemeManager 创建主题管理器实例
func NewThemeManager() *ThemeManagerImpl {
	return &ThemeManagerImpl{
		themes: make(map[string]*ThemeInfo),
	}
}

// SwitchTheme 切换主题
func (m *ThemeManagerImpl) SwitchTheme(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	newTheme, exists := m.themes[id]
	if !exists {
		// 尝试动态加载主题
		cfgs, err := configs.GetConfig()
		if err != nil {
			return fmt.Errorf("failed to get config: %w", err)
		}

		// 扫描所有目录查找匹配的主题ID
		var themePath string
		var themeJson []byte

		entries, err := os.ReadDir(cfgs.ThemeConfig.ThemeDir)
		if err != nil {
			return fmt.Errorf("failed to scan theme directory: %w", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			dirPath := filepath.Join(cfgs.ThemeConfig.ThemeDir, entry.Name())
			configPath := filepath.Join(dirPath, cfgs.ThemeConfig.ThemeConfigFile)

			if configData, err := os.ReadFile(configPath); err == nil {
				var tempInfo ThemeInfo
				if err := json.Unmarshal(configData, &tempInfo); err == nil && tempInfo.ID == id {
					themePath = dirPath
					themeJson = configData
					break
				}
			}
		}

		if themePath == "" {
			return fmt.Errorf("theme %s not found", id)
		}

		var themeConfig ThemeInfo
		if err := json.Unmarshal(themeJson, &themeConfig); err != nil {
			return fmt.Errorf("invalid theme config for %s: %w", id, err)
		}

		themeConfig.Path = themePath
		themeConfig.Status = ""
		themeConfig.IsActive = false

		// 编译主题
		if err := theme.ExecuteBuildScript(themePath); err != nil {
			return fmt.Errorf("failed to build theme %s: %w", id, err)
		}

		// 将新加载的主题添加到内存中
		m.themes[themeConfig.ID] = &themeConfig
		newTheme = &themeConfig
	}

	// 停用当前激活的主题
	if m.activeTheme != "" {
		if currentTheme, ok := m.themes[m.activeTheme]; ok {
			currentTheme.IsActive = false
			currentTheme.Status = ""
		}
	}

	// 激活新主题
	newTheme.IsActive = true
	newTheme.Status = consts.ThemeStatusActive
	newTheme.LoadedAt = time.Now().Unix()
	m.activeTheme = id

	// 持久化到配置文件
	if err := configs.UpdateField(func(config *configs.Config) {
		config.ThemeConfig.LastActiveTheme = id
	}); err != nil {
		global.SysLog.Warnf("Failed to persist active theme to config: %v", err)
	}

	global.SysLog.Infof("Theme switched: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
		newTheme.ID, newTheme.Name, newTheme.Version, newTheme.Repository, newTheme.Author, newTheme.Path, newTheme.Status)
	return nil
}

// GetActiveTheme 获取当前激活的主题
func (m *ThemeManagerImpl) GetActiveTheme() (*ThemeInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	theme, exists := m.themes[m.activeTheme]
	if !exists {
		return nil, fmt.Errorf("active theme %s not found", m.activeTheme)
	}

	// 返回主题信息拷贝
	copy := *theme
	return &copy, nil
}

// ListThemes 列出所有主题
func (m *ThemeManagerImpl) ListThemes() ([]*ThemeInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	cfgs, err := configs.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	entries, err := os.ReadDir(cfgs.ThemeConfig.ThemeDir)
	if err != nil {
		return nil, fmt.Errorf("failed to scan theme directory: %w", err)
	}

	var themes []*ThemeInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		themeID := entry.Name()
		themePath := filepath.Join(cfgs.ThemeConfig.ThemeDir, themeID)

		// 尝试读取主题配置文件获取真实ID
		configPath := filepath.Join(themePath, cfgs.ThemeConfig.ThemeConfigFile)
		if configData, err := os.ReadFile(configPath); err == nil {
			var themeInfo ThemeInfo
			if err := json.Unmarshal(configData, &themeInfo); err == nil {
				// 使用配置文件中的真实ID
				realID := themeInfo.ID

				// 如果主题已加载，使用已加载的信息
				if loadedTheme, exists := m.themes[realID]; exists {
					themes = append(themes, loadedTheme)
				} else {
					themeInfo.Path = themePath
					themeInfo.Status = ""
					themeInfo.IsActive = (realID == m.activeTheme)
					themes = append(themes, &themeInfo)
				}
			}
		}
	}

	return themes, nil
}

// InitializeTheme 初始化主题系统，加载上次激活的主题或默认主题
func (m *ThemeManagerImpl) InitializeTheme() error {
	cfgs, err := configs.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	targetThemeID := cfgs.ThemeConfig.LastActiveTheme
	if targetThemeID == "" {
		targetThemeID = cfgs.ThemeConfig.DefaultTheme
	}

	// 扫描所有目录查找匹配的主题ID
	var themePath string
	var themeJson []byte

	entries, err := os.ReadDir(cfgs.ThemeConfig.ThemeDir)
	if err != nil {
		return fmt.Errorf("failed to scan theme directory: %w", err)
	}

	// 定义扫描函数，避免重复代码
	findTheme := func(themeID string) (string, []byte) {
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			dirPath := filepath.Join(cfgs.ThemeConfig.ThemeDir, entry.Name())
			configPath := filepath.Join(dirPath, cfgs.ThemeConfig.ThemeConfigFile)

			if configData, err := os.ReadFile(configPath); err == nil {
				var tempInfo ThemeInfo
				if err := json.Unmarshal(configData, &tempInfo); err == nil && tempInfo.ID == themeID {
					return dirPath, configData
				}
			}
		}
		return "", nil
	}

	// 查找目标主题
	themePath, themeJson = findTheme(targetThemeID)

	// 如果没找到目标主题，尝试回退到默认主题
	if themePath == "" && targetThemeID != cfgs.ThemeConfig.DefaultTheme {
		global.SysLog.Warnf("Failed to load theme %s, falling back to default theme %s", targetThemeID, cfgs.ThemeConfig.DefaultTheme)
		targetThemeID = cfgs.ThemeConfig.DefaultTheme
		themePath, themeJson = findTheme(targetThemeID)
	}

	if themePath == "" {
		return fmt.Errorf("theme %s not found", targetThemeID)
	}

	var themeConfig ThemeInfo
	if err := json.Unmarshal(themeJson, &themeConfig); err != nil {
		return fmt.Errorf("invalid theme config for %s: %w", targetThemeID, err)
	}

	themeConfig.Path = themePath

	// 编译主题
	if err := theme.ExecuteBuildScript(themePath); err != nil {
		return fmt.Errorf("failed to build theme %s: %w", targetThemeID, err)
	}

	themeConfig.Status = consts.ThemeStatusActive
	themeConfig.LoadedAt = time.Now().Unix()
	themeConfig.IsActive = true

	m.themes[themeConfig.ID] = &themeConfig
	m.activeTheme = themeConfig.ID

	if err := configs.UpdateField(func(config *configs.Config) {
		config.ThemeConfig.LastActiveTheme = themeConfig.ID
	}); err != nil {
		global.SysLog.Warnf("Failed to persist active theme to config: %v", err)
	}

	global.SysLog.Infof("Theme system initialized: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
		themeConfig.ID, themeConfig.Name, themeConfig.Version, themeConfig.Repository, themeConfig.Author, themeConfig.Path, themeConfig.Status)
	return nil
}

// Shutdown 关闭主题系统
func (m *ThemeManagerImpl) Shutdown() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, theme := range m.themes {
		global.SysLog.Infof("Theme shutdown: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
			theme.ID, theme.Name, theme.Version, theme.Repository, theme.Author, theme.Path, theme.Status)
	}

	m.themes = make(map[string]*ThemeInfo)
	m.activeTheme = ""
}
