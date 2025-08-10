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
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrPluginParamInvalid)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrPluginParamInvalid)))
		return
	}

	response, err := tc.themeService.SwitchTheme(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPluginSystemError, errorx.KV("error", err.Error()))))
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
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPluginNotFound)))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListThemes 列举主题
// @Router /api/theme/list [get]
func (tc *ThemeController) ListThemes(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListThemesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrPluginParamInvalid)))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrPluginParamInvalid)))
		return
	}

	response, err := tc.themeService.ListThemes(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPluginSystemError, errorx.KV("error", err.Error()))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
