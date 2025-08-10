package impl

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/theme"
	"github.com/Done-0/jank/internal/theme/impl"
	themeUtils "github.com/Done-0/jank/internal/utils/theme"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

// ThemeServiceImpl 主题服务实现
type ThemeServiceImpl struct{}

// NewThemeService 创建主题服务实例
func NewThemeService() service.ThemeService {
	return &ThemeServiceImpl{}
}

// SwitchTheme 切换主题
func (s *ThemeServiceImpl) SwitchTheme(c *app.RequestContext, req *dto.SwitchThemeRequest) (*vo.SwitchThemeResponse, error) {
	if req.Rebuild {
		themes, err := theme.GlobalThemeManager.ListThemes()
		if err != nil {
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
			return nil, fmt.Errorf("theme %s not found", req.ID)
		}

		// 执行构建
		if err := themeUtils.ExecuteBuildScript(targetTheme.Path); err != nil {
			return nil, fmt.Errorf("failed to rebuild theme %s: %w", req.ID, err)
		}
	}

	if err := theme.GlobalThemeManager.SwitchTheme(req.ID); err != nil {
		return nil, fmt.Errorf("failed to switch theme: %w", err)
	}

	return &vo.SwitchThemeResponse{
		Message: fmt.Sprintf("Theme switched to %s successfully", req.ID),
	}, nil
}

// GetActiveTheme 获取当前激活的主题
func (s *ThemeServiceImpl) GetActiveTheme(c *app.RequestContext) (*vo.GetActiveThemeResponse, error) {
	themeInfo, err := theme.GlobalThemeManager.GetActiveTheme()
	if err != nil {
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

// ListThemes 列举主题
func (s *ThemeServiceImpl) ListThemes(c *app.RequestContext, req *dto.ListThemesRequest) (*vo.ListThemesResponse, error) {
	discoveredThemes, err := theme.GlobalThemeManager.ListThemes()
	if err != nil {
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
