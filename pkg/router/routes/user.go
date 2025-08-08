// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
)

// RegisterUserRoutes 注册用户相关路由
// 参数：
//
//	r: Hertz 路由组数组，r[0] 为 API v1 版本组
func RegisterUserRoutes(r ...*route.RouterGroup) {
	// api v1 group
	apiV1 := r[0]
	userGroup := apiV1.Group("/user")

	// TODO: 实现具体的路由处理函数后，取消注释以下路由
	// userGroup.POST("/register", account.RegisterAcc)
	// userGroup.POST("/login", account.LoginAccount)
	// userGroup.GET("/profile", account.GetAccount, middleware.AuthMiddleware())
	// userGroup.POST("/update", account.UpdateAccount, middleware.AuthMiddleware())
	// userGroup.POST("/logout", account.LogoutAccount, middleware.AuthMiddleware())
	// userGroup.POST("/resetPassword", account.ResetPassword, middleware.AuthMiddleware())

	_ = userGroup // 避免未使用变量警告
}
