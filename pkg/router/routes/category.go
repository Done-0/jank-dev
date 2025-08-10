// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
)

// RegisterCategoryRoutes 注册分类相关路由
// 参数：
//
//	r: Hertz 路由组数组，r[0] 为 API v1 版本组
func RegisterCategoryRoutes(r ...*route.RouterGroup) {
	// api v1 group
	apiV1 := r[0]
	categoryGroupV1 := apiV1.Group("/category")

	// TODO: 实现具体的路由处理函数后，取消注释以下路由
	// categoryGroupV1.GET("/getOne", category.GetOneCategory)
	// categoryGroupV1.GET("/getTree", category.GetCategoryTree)
	// categoryGroupV1.GET("/getChildrenTree", category.GetCategoryChildrenTree)
	// categoryGroupV1.POST("/create", category.CreateOneCategory, middleware.AuthMiddleware())
	// categoryGroupV1.POST("/update", category.UpdateOneCategory, middleware.AuthMiddleware())
	// categoryGroupV1.POST("/delete", category.DeleteOneCategory, middleware.AuthMiddleware())

	_ = categoryGroupV1 // 避免未使用变量警告
}
