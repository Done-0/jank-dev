// Package controller 主题控制器
// 创建者：Done-0
// 创建时间：2025-08-05
package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
	"github.com/Done-0/jank/internal/utils/validator"
	"github.com/Done-0/jank/internal/utils/vo"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
)

// ThemeController 主题控制器
type ThemeController struct {
	themeService service.ThemeService
}

// NewThemeController 创建主题控制器
func NewThemeController(themeService service.ThemeService) *ThemeController {
	return &ThemeController{
		themeService: themeService,
	}
}

// SwitchTheme 切换主题
// @Router /api/theme/switch [post]
func (tc *ThemeController) SwitchTheme(ctx context.Context, c *app.RequestContext) {
	req := new(dto.SwitchThemeRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "request_body"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := tc.themeService.SwitchTheme(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "switch theme failed"))))
		return
	}

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetActiveTheme 获取当前激活主题
// @Router /api/theme/get [get]
func (tc *ThemeController) GetActiveTheme(ctx context.Context, c *app.RequestContext) {
	response, err := tc.themeService.GetActiveTheme(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "get active theme failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListThemes 列举主题
// @Router /api/theme/list [get]
func (tc *ThemeController) ListThemes(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListThemesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "query_params"), errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := tc.themeService.ListThemes(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "list themes failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ServeHomePage 提供主题首页
func (tc *ThemeController) ServeHomePage(ctx context.Context, c *app.RequestContext) {
	homePagePath, err := tc.themeService.ServeHomePage(c)
	if err != nil {
		c.AbortWithStatus(consts.StatusInternalServerError)
		return
	}

	c.File(homePagePath)
}

// ServeStaticResource 提供静态资源文件
func (tc *ThemeController) ServeStaticResource(ctx context.Context, c *app.RequestContext) {
	staticResourcePath, err := tc.themeService.ServeStaticResource(c, string(c.Path()))
	if err != nil {
		c.AbortWithStatus(consts.StatusNotFound)
		return
	}

	c.File(staticResourcePath)
}
