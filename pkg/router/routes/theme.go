package routes

import (
	"context"
	"log"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/internal/theme"
	"github.com/Done-0/jank/pkg/wire"
)

// RegisterThemeFrontendRoutes 注册主题前端路由（动态处理器）
func RegisterThemeFrontendRoutes(h *server.Hertz) {
	// 动态首页处理器
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		activeTheme, err := theme.GlobalThemeManager.GetActiveTheme()
		if err != nil {
			log.Printf("Failed to get active theme: %v", err)
			c.AbortWithStatus(500)
			return
		}

		c.File(filepath.Join(activeTheme.Path, activeTheme.IndexFilePath))
	})

	// 通用静态资源处理器 - 处理所有未匹配的路径
	h.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		activeTheme, err := theme.GlobalThemeManager.GetActiveTheme()
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		// 获取请求的文件路径（去掉开头的 /）
		requestedFile := strings.TrimPrefix(string(c.Path()), "/")
		
		// 获取主题的构建目录（如 dist/）
		buildDir := filepath.Dir(activeTheme.IndexFilePath)
		
		// 拼接完整文件路径：themes/主题名/构建目录/请求文件
		fullPath := filepath.Join(activeTheme.Path, buildDir, requestedFile)
		
		c.File(fullPath)
	})
}

// RegisterThemeRoutes 注册主题路由
func RegisterThemeRoutes(r *route.RouterGroup) {
	themeController, err := wire.NewThemeController()
	if err != nil {
		log.Fatalf("Failed to initialize theme controller: %v", err)
	}

	// 主题路由组
	themeGroup := r.Group("/theme")
	{
		// POST 方法
		themeGroup.POST("/switch", themeController.SwitchTheme) // 切换主题

		// GET 方法
		themeGroup.GET("/get", themeController.GetActiveTheme) // 获取当前激活主题
		themeGroup.GET("/list", themeController.ListThemes)    // 列举所有主题
	}
}
