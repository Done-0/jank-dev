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

		themePath := filepath.Join(cfgs.ThemeConfig.ThemeDir, id)
		configPath := filepath.Join(themePath, cfgs.ThemeConfig.ThemeConfigFile)

		themeJson, err := os.ReadFile(configPath)
		if err != nil {
			return fmt.Errorf("theme %s not found", id)
		}

		var themeConfig ThemeInfo
		if err := json.Unmarshal(themeJson, &themeConfig); err != nil {
			return fmt.Errorf("invalid theme config for %s: %w", id, err)
		}

		themeConfig.Path = themePath
		themeConfig.Status = consts.ThemeStatusReady
		themeConfig.IsActive = false

		// 将新加载的主题添加到内存中
		m.themes[themeConfig.ID] = &themeConfig
		newTheme = &themeConfig
	}

	// 停用当前激活的主题
	if m.activeTheme != "" {
		if currentTheme, ok := m.themes[m.activeTheme]; ok {
			currentTheme.IsActive = false
			currentTheme.Status = consts.ThemeStatusReady
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

		// 如果主题已加载，使用已加载的信息，否则尝试加载配置文件
		if loadedTheme, exists := m.themes[themeID]; exists {
			themes = append(themes, loadedTheme)
		} else {
			// 尝试读取主题配置文件
			configPath := filepath.Join(themePath, cfgs.ThemeConfig.ThemeConfigFile)
			if themeJson, err := os.ReadFile(configPath); err == nil {
				var themeInfo ThemeInfo
				if err := json.Unmarshal(themeJson, &themeInfo); err == nil {
					themeInfo.Path = themePath
					themeInfo.Status = consts.ThemeStatusReady
					themeInfo.IsActive = (themeID == m.activeTheme)
					themes = append(themes, &themeInfo)
				} else {
					// 配置文件解析失败，使用基本信息
					themeInfo := &ThemeInfo{
						ID:       themeID,
						Name:     themeID,
						Path:     themePath,
						Status:   consts.ThemeStatusReady,
						IsActive: false,
					}
					themes = append(themes, themeInfo)
				}
			} else {
				// 配置文件不存在，使用基本信息
				themeInfo := &ThemeInfo{
					ID:       themeID,
					Name:     themeID,
					Path:     themePath,
					Status:   consts.ThemeStatusReady,
					IsActive: false,
				}
				themes = append(themes, themeInfo)
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

	// 尝试加载目标主题，如果失败则回退到默认主题
	themePath := filepath.Join(cfgs.ThemeConfig.ThemeDir, targetThemeID)
	themeJson, err := os.ReadFile(filepath.Join(themePath, cfgs.ThemeConfig.ThemeConfigFile))
	if err != nil {
		if targetThemeID != cfgs.ThemeConfig.DefaultTheme {
			global.SysLog.Warnf("Failed to load theme %s, falling back to default theme %s: %v", targetThemeID, cfgs.ThemeConfig.DefaultTheme, err)
			targetThemeID = cfgs.ThemeConfig.DefaultTheme
			themePath = filepath.Join(cfgs.ThemeConfig.ThemeDir, targetThemeID)
			themeJson, err = os.ReadFile(filepath.Join(themePath, cfgs.ThemeConfig.ThemeConfigFile))
			if err != nil {
				return fmt.Errorf("failed to read default theme config for %s: %w", targetThemeID, err)
			}
		} else {
			return fmt.Errorf("failed to read theme config for %s: %w", targetThemeID, err)
		}
	}

	var themeConfig ThemeInfo
	if err := json.Unmarshal(themeJson, &themeConfig); err != nil {
		return fmt.Errorf("invalid theme config for %s: %w", targetThemeID, err)
	}

	themeConfig.Path = themePath

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
