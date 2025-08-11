package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// ThemeService 主题服务接口
type ThemeService interface {
	SwitchTheme(c *app.RequestContext, req *dto.SwitchThemeRequest) (*vo.SwitchThemeResponse, error)
	GetActiveTheme(c *app.RequestContext) (*vo.GetActiveThemeResponse, error)
	ListThemes(c *app.RequestContext, req *dto.ListThemesRequest) (*vo.ListThemesResponse, error)
	
	// 前端路由服务方法
	ServeHomePage(c *app.RequestContext) (string, error)
	ServeStaticResource(c *app.RequestContext, requestPath string) (string, error)
}
