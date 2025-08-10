// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
)

// RegisterCommentRoutes 注册评论相关路由
// 参数：
//
//	r: Hertz 路由组数组，r[0] 为 API v1 版本组
func RegisterCommentRoutes(r ...*route.RouterGroup) {
	// api v1 group
	apiV1 := r[0]
	commentGroup := apiV1.Group("/comment")

	// TODO: 实现具体的路由处理函数后，取消注释以下路由
	// commentGroup.GET("/getOne", comment.GetOneComment)
	// commentGroup.GET("/getGraph", comment.GetCommentGraph)
	// commentGroup.POST("/create", comment.CreateOneComment, middleware.AuthMiddleware())
	// commentGroup.POST("/delete", comment.DeleteOneComment, middleware.AuthMiddleware())

	_ = commentGroup // 避免未使用变量警告
}
