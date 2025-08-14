package impl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/theme"
	"github.com/Done-0/jank/internal/theme/impl"
	"github.com/Done-0/jank/internal/types/consts"
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
	// 获取所有可用主题列表
	availableThemes, err := theme.GlobalThemeManager.ListThemes()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list available themes: %v", err)
		return nil, fmt.Errorf("failed to list available themes: %w", err)
	}

	// 在可用主题中查找目标主题
	var requestedTheme *impl.ThemeInfo
	for _, availableTheme := range availableThemes {
		isTargetTheme := availableTheme.ID == req.ID
		switch {
		case isTargetTheme:
			requestedTheme = availableTheme
		}

		// 如果找到目标主题，退出循环
		if requestedTheme != nil {
			break
		}
	}

	// 确保目标主题存在
	switch requestedTheme {
	case nil:
		logger.BizLogger(c).Errorf("requested theme does not exist: %s", req.ID)
		return nil, fmt.Errorf("theme %s not found", req.ID)
	}

	// 验证主题类型是否匹配请求
	requestedThemeType := req.ThemeType
	actualThemeType := requestedTheme.Type

	switch {
	case requestedThemeType != actualThemeType:
		logger.BizLogger(c).Errorf("theme type mismatch - requested: %s, actual: %s",
			requestedThemeType, actualThemeType)
		return nil, fmt.Errorf("theme type mismatch: theme %s is type %s, not %s",
			requestedTheme.ID, actualThemeType, requestedThemeType)
	}

	// 验证请求来源是否有权限切换此类型主题
	currentRequestPath := string(c.Path())
	isRequestFromConsole := strings.HasPrefix(currentRequestPath, "/console")
	isRequestingFrontendTheme := actualThemeType == consts.ThemeTypeFrontend

	switch {
	case !isRequestFromConsole && !isRequestingFrontendTheme:
		logger.BizLogger(c).Errorf("frontend page attempted to switch to non-frontend theme: %s", actualThemeType)
		return nil, fmt.Errorf("frontend pages can only switch to frontend themes")
	}

	// 执行主题构建（如果需要重新构建）
	switch {
	case req.Rebuild:
		themeSourcePath := requestedTheme.Path
		switch buildErr := themeUtils.ExecuteBuildScript(themeSourcePath); buildErr {
		case nil:
			logger.BizLogger(c).Infof("theme %s rebuilt successfully from path: %s", req.ID, themeSourcePath)
		default:
			logger.BizLogger(c).Errorf("failed to rebuild theme %s: %v", req.ID, buildErr)
			return nil, fmt.Errorf("failed to rebuild theme %s: %w", req.ID, buildErr)
		}
	}

	// 执行主题切换操作
	switch switchErr := theme.GlobalThemeManager.SwitchThemeByType(req.ID, req.ThemeType); switchErr {
	case nil:
		logger.BizLogger(c).Infof("theme switched successfully: %s (type: %s)", req.ID, req.ThemeType)
	default:
		logger.BizLogger(c).Errorf("failed to switch theme: %v", switchErr)
		return nil, fmt.Errorf("failed to switch theme: %w", switchErr)
	}

	return &vo.SwitchThemeResponse{
		Message: fmt.Sprintf("%s theme switched to %s successfully", req.ThemeType, req.ID),
	}, nil
}

// GetActiveTheme 获取当前激活的主题
func (s *ThemeServiceImpl) GetActiveTheme(c *app.RequestContext) (*vo.GetActiveThemeResponse, error) {
	// 根据请求路径确定主题类型
	var themeType string
	if strings.HasPrefix(string(c.Path()), "/console") {
		themeType = consts.ThemeTypeConsole
	} else {
		themeType = consts.ThemeTypeFrontend
	}

	themeInfo, err := theme.GlobalThemeManager.GetActiveThemeByType(themeType)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get active %s theme: %v", themeType, err)
		return nil, fmt.Errorf("failed to get active %s theme: %w", themeType, err)
	}

	themeResponse := &vo.GetThemeResponse{
		ID:            themeInfo.ID,
		Name:          themeInfo.Name,
		Version:       themeInfo.Version,
		Author:        themeInfo.Author,
		Description:   themeInfo.Description,
		Repository:    themeInfo.Repository,
		Preview:       themeInfo.Preview,
		Type:          themeInfo.Type,
		IndexFilePath: themeInfo.IndexFilePath,
		StaticDirPath: themeInfo.StaticDirPath,
		Status:        themeInfo.Status,
		LoadedAt:      themeInfo.LoadedAt,
		Path:          themeInfo.Path,
		IsActive:      themeInfo.IsActive,
	}

	return &vo.GetActiveThemeResponse{
		Theme: themeResponse,
	}, nil
}

// ListThemes 列举主题逻辑
func (s *ThemeServiceImpl) ListThemes(c *app.RequestContext, req *dto.ListThemesRequest) (*vo.ListThemesResponse, error) {
	discoveredThemes, err := theme.GlobalThemeManager.ListThemes()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list themes: %v", err)
		return &vo.ListThemesResponse{}, fmt.Errorf("failed to list themes: %w", err)
	}

	// 过滤和转换主题
	var filteredThemes []vo.GetThemeResponse
	for _, themeInfo := range discoveredThemes {
		if req.Status != "" && themeInfo.Status != req.Status {
			continue
		}
		filteredThemes = append(filteredThemes, vo.GetThemeResponse{
			ID:            themeInfo.ID,
			Name:          themeInfo.Name,
			Version:       themeInfo.Version,
			Author:        themeInfo.Author,
			Description:   themeInfo.Description,
			Repository:    themeInfo.Repository,
			Preview:       themeInfo.Preview,
			Type:          themeInfo.Type,
			IndexFilePath: themeInfo.IndexFilePath,
			StaticDirPath: themeInfo.StaticDirPath,
			Status:        themeInfo.Status,
			LoadedAt:      themeInfo.LoadedAt,
			Path:          themeInfo.Path,
			IsActive:      themeInfo.IsActive,
		})
	}

	// 分页处理
	total := int64(len(filteredThemes))
	start := (req.PageNo - 1) * req.PageSize
	end := start + req.PageSize

	var pagedThemes []vo.GetThemeResponse
	if start >= total {
		pagedThemes = []vo.GetThemeResponse{}
	} else {
		if end > total {
			end = total
		}
		pagedThemes = filteredThemes[start:end]
	}

	logger.BizLogger(c).Infof("listed %d themes (total: %d, page: %d, size: %d)", len(pagedThemes), total, req.PageNo, req.PageSize)
	return &vo.ListThemesResponse{
		List:  pagedThemes,
		Total: total,
	}, nil
}

// ServeHomePage 获取主题首页文件路径逻辑
func (s *ThemeServiceImpl) ServeHomePage(c *app.RequestContext) (string, error) {
	// 根据请求路径确定主题类型
	var themeType string
	if strings.HasPrefix(string(c.Path()), "/console") {
		themeType = consts.ThemeTypeConsole
	} else {
		themeType = consts.ThemeTypeFrontend
	}

	activeTheme, err := theme.GlobalThemeManager.GetActiveThemeByType(themeType)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get active %s theme: %v", themeType, err)
		return "", fmt.Errorf("failed to get active %s theme: %w", themeType, err)
	}

	logger.BizLogger(c).Infof("serving home page for %s theme: %s", themeType, activeTheme.ID)
	return filepath.Join(activeTheme.Path, activeTheme.IndexFilePath), nil
}

// ServeStaticResource 获取静态资源文件路径逻辑
func (s *ThemeServiceImpl) ServeStaticResource(c *app.RequestContext, requestPath string) (string, error) {
	// 跳过API路径
	if strings.HasPrefix(requestPath, "/api/") {
		logger.BizLogger(c).Warnf("API path serving attempted: %s", requestPath)
		return "", fmt.Errorf("API paths are not served as static resources")
	}

	// 确定主题类型（考虑Referer头）
	var themeType string
	if strings.HasPrefix(requestPath, "/console") {
		themeType = consts.ThemeTypeConsole
	} else {
		// 检查 Refere r头，如果来自 console 页面，则使用 console 主题
		referer := string(c.GetHeader("Referer"))
		if strings.Contains(referer, "/console") {
			themeType = consts.ThemeTypeConsole
		} else {
			themeType = consts.ThemeTypeFrontend
		}
	}

	activeTheme, err := theme.GlobalThemeManager.GetActiveThemeByType(themeType)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get active %s theme: %v", themeType, err)
		return "", fmt.Errorf("failed to get active %s theme: %w", themeType, err)
	}

	// 路径安全验证
	cleanedPath := filepath.Clean(strings.TrimPrefix(requestPath, "/"))
	if strings.Contains(cleanedPath, "..") || strings.Contains(cleanedPath, "\\") {
		logger.BizLogger(c).Errorf("path traversal attempt detected: %s", requestPath)
		return "", fmt.Errorf("invalid file path: path traversal not allowed")
	}

	// 构建资源文件的完整路径
	themeBuildDir := filepath.Dir(activeTheme.IndexFilePath)
	themeBasePath := filepath.Join(activeTheme.Path, themeBuildDir)
	resourcePath := filepath.Join(themeBasePath, cleanedPath)

	absoluteResourcePath, err := filepath.Abs(resourcePath)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to resolve absolute path for resource: %v", err)
		return "", fmt.Errorf("failed to resolve resource path")
	}

	// 安全边界检查
	appConfig, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to load application config: %v", err)
		return "", fmt.Errorf("failed to load configuration")
	}

	allowedThemeDir, err := filepath.Abs(appConfig.ThemeConfig.ThemeDir)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to resolve theme directory: %v", err)
		return "", fmt.Errorf("failed to resolve theme directory")
	}

	if !strings.HasPrefix(absoluteResourcePath, allowedThemeDir) {
		logger.BizLogger(c).Errorf("security violation: resource %s is outside theme directory", absoluteResourcePath)
		return "", fmt.Errorf("access denied: resource outside allowed directory")
	}

	// 检查文件是否存在，处理SPA路由回退
	if _, err := os.Stat(absoluteResourcePath); os.IsNotExist(err) {
		// 判断是否为SPA路由（无文件扩展名）
		isSPARoute := !strings.Contains(cleanedPath, ".")
		if isSPARoute {
			logger.BizLogger(c).Infof("SPA route detected: %s, serving index.html for %s theme", requestPath, themeType)
			return filepath.Join(activeTheme.Path, activeTheme.IndexFilePath), nil
		}

		// 静态资源不存在
		logger.BizLogger(c).Warnf("static resource not found: %s", absoluteResourcePath)
		return "", fmt.Errorf("static resource not found: %s", cleanedPath)
	}

	return absoluteResourcePath, nil
}
