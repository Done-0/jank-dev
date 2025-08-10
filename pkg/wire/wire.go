//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/Done-0/jank/pkg/serve/controller"
)

// NewPluginController 使用 Wire 初始化插件控制器
func NewPluginController() (*controller.PluginController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewPluginController,
	))
}

// NewRBACController 使用 Wire 初始化RBAC控制器
func NewRBACController() (*controller.RBACController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewRBACController,
	))
}

// NewThemeController 使用 Wire 初始化主题控制器
func NewThemeController() (*controller.ThemeController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewThemeController,
	))
}
