package impl

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/theme"
	"github.com/Done-0/jank/internal/theme/impl"
	"github.com/Done-0/jank/internal/utils/logger"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"

	themeUtils "github.com/Done-0/jank/internal/utils/theme"
)

// ThemeServiceImpl 主题服务实现
type ThemeServiceImpl struct{}

// NewThemeService 创建主题服务实例
func NewThemeService() service.ThemeService {
	return &ThemeServiceImpl{}
}

// SwitchTheme 切换主题逻辑
func (s *ThemeServiceImpl) SwitchTheme(c *app.RequestContext, req *dto.SwitchThemeRequest) (*vo.SwitchThemeResponse, error) {
	if req.Rebuild {
		themes, err := theme.GlobalThemeManager.ListThemes()
		if err != nil {
			logger.BizLogger(c).Errorf("failed to list themes: %v", err)
			return nil, fmt.Errorf("failed to list themes: %w", err)
		}

		var targetTheme *impl.ThemeInfo
		for _, t := range themes {
			if t.ID == req.ID {
				targetTheme = t
				break
			}
		}

		if targetTheme == nil {
			logger.BizLogger(c).Errorf("theme not found: %s", req.ID)
			return nil, fmt.Errorf("theme %s not found", req.ID)
		}

		if err := themeUtils.ExecuteBuildScript(targetTheme.Path); err != nil {
			logger.BizLogger(c).Errorf("failed to rebuild theme %s: %v", req.ID, err)
			return nil, fmt.Errorf("failed to rebuild theme %s: %w", req.ID, err)
		}
	}

	if err := theme.GlobalThemeManager.SwitchTheme(req.ID); err != nil {
		logger.BizLogger(c).Errorf("failed to switch theme to %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to switch theme: %w", err)
	}

	return &vo.SwitchThemeResponse{
		Message: fmt.Sprintf("Theme switched to %s successfully", req.ID),
	}, nil
}

// GetActiveTheme 获取当前激活的主题逻辑
func (s *ThemeServiceImpl) GetActiveTheme(c *app.RequestContext) (*vo.GetActiveThemeResponse, error) {
	themeInfo, err := theme.GlobalThemeManager.GetActiveTheme()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get active theme: %v", err)
		return nil, fmt.Errorf("failed to get active theme: %w", err)
	}

	return &vo.GetActiveThemeResponse{
		Theme: &vo.GetThemeResponse{
			ID:            themeInfo.ID,
			Name:          themeInfo.Name,
			Version:       themeInfo.Version,
			Author:        themeInfo.Author,
			Description:   themeInfo.Description,
			Repository:    themeInfo.Repository,
			Preview:       themeInfo.Preview,
			IndexFilePath: themeInfo.IndexFilePath,
			StaticDirPath: themeInfo.StaticDirPath,
			Status:        themeInfo.Status,
			LoadedAt:      themeInfo.LoadedAt,
			Path:          themeInfo.Path,
			IsActive:      themeInfo.IsActive,
		},
	}, nil
}

// ListThemes 列举主题逻辑
func (s *ThemeServiceImpl) ListThemes(c *app.RequestContext, req *dto.ListThemesRequest) (*vo.ListThemesResponse, error) {
	discoveredThemes, err := theme.GlobalThemeManager.ListThemes()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list themes: %v", err)
		return &vo.ListThemesResponse{}, fmt.Errorf("failed to list themes: %w", err)
	}

	var filteredThemes []vo.GetThemeResponse
	for _, themeInfo := range discoveredThemes {
		if req.Status != "" && themeInfo.Status != req.Status {
			continue
		}

		discoveryVO := vo.GetThemeResponse{
			ID:            themeInfo.ID,
			Name:          themeInfo.Name,
			Version:       themeInfo.Version,
			Author:        themeInfo.Author,
			Description:   themeInfo.Description,
			Repository:    themeInfo.Repository,
			Preview:       themeInfo.Preview,
			IndexFilePath: themeInfo.IndexFilePath,
			StaticDirPath: themeInfo.StaticDirPath,
			Status:        themeInfo.Status,
			LoadedAt:      themeInfo.LoadedAt,
			Path:          themeInfo.Path,
			IsActive:      themeInfo.IsActive,
		}
		filteredThemes = append(filteredThemes, discoveryVO)
	}

	total := int64(len(filteredThemes))

	pageNo := req.PageNo
	pageSize := req.PageSize
	start := (pageNo - 1) * pageSize
	end := start + pageSize
	if start >= total {
		filteredThemes = []vo.GetThemeResponse{}
	} else {
		if end > total {
			end = total
		}
		filteredThemes = filteredThemes[start:end]
	}

	return &vo.ListThemesResponse{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		List:     filteredThemes,
	}, nil
}

// ServeHomePage 获取主题首页文件路径逻辑
func (s *ThemeServiceImpl) ServeHomePage(c *app.RequestContext) (string, error) {
	activeTheme, err := theme.GlobalThemeManager.GetActiveTheme()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get active theme for home page: %v", err)
		return "", fmt.Errorf("failed to get active theme: %w", err)
	}

	return filepath.Join(activeTheme.Path, activeTheme.IndexFilePath), nil
}

// ServeStaticResource 获取静态资源文件路径逻辑
func (s *ThemeServiceImpl) ServeStaticResource(c *app.RequestContext, requestPath string) (string, error) {
	// 跳过 API 路径
	if strings.HasPrefix(requestPath, "/api/") {
		logger.BizLogger(c).Warnf("attempted to serve API path as static resource: %s", requestPath)
		return "", fmt.Errorf("API paths are not served as static resources")
	}

	activeTheme, err := theme.GlobalThemeManager.GetActiveTheme()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get active theme for static resource: %v", err)
		return "", fmt.Errorf("failed to get active theme: %w", err)
	}

	// 安全检查：防止路径遍历攻击
	requestedFile := filepath.Clean(strings.TrimPrefix(requestPath, "/"))
	if strings.Contains(requestedFile, "..") || strings.Contains(requestedFile, "\\") {
		logger.BizLogger(c).Errorf("path traversal attempt detected: %s", requestPath)
		return "", fmt.Errorf("invalid file path: path traversal not allowed")
	}

	buildDir := filepath.Dir(activeTheme.IndexFilePath)
	basePath := filepath.Join(activeTheme.Path, buildDir)
	fullPath := filepath.Join(basePath, requestedFile)

	absFullPath, err := filepath.Abs(fullPath)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to resolve file path: %v", err)
		return "", fmt.Errorf("failed to resolve file path")
	}

	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return "", fmt.Errorf("failed to get configuration")
	}
	absThemesDir, err := filepath.Abs(cfgs.ThemeConfig.ThemeDir)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to resolve themes directory: %v", err)
		return "", fmt.Errorf("failed to resolve themes directory")
	}

	// 安全边界检查
	if !strings.HasPrefix(absFullPath, absThemesDir) {
		logger.BizLogger(c).Errorf("access denied: %s outside themes directory", absFullPath)
		return "", fmt.Errorf("access denied: file outside themes directory")
	}

	return fullPath, nil
}
