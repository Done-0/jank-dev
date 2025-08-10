// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
)

// RegisterPostRoutes 注册文章相关路由
// 参数：
//
//	r: Hertz 路由组数组，r[0] 为 API v1 版本组
func RegisterPostRoutes(r ...*route.RouterGroup) {
	// api v1 group
	apiV1 := r[0]
	postGroup := apiV1.Group("/post")

	// TODO: 实现具体的路由处理函数后，取消注释以下路由
	// postGroup.GET("/getOne", post.GetOnePost)
	// postGroup.GET("/getAll", post.GetAllPosts)
	// postGroup.POST("/create", post.CreateOnePost, middleware.AuthMiddleware())
	// postGroup.POST("/update", post.UpdateOnePost, middleware.AuthMiddleware())
	// postGroup.POST("/delete", post.DeleteOnePost, middleware.AuthMiddleware())

	_ = postGroup // 避免未使用变量警告
}
