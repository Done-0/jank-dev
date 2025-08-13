// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-13
package routes

import (
	"log"

	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/pkg/wire"
)

// RegisterCategoryRoutes 注册分类相关路由
func RegisterCategoryRoutes(r *route.RouterGroup) {
	categoryController, err := wire.NewCategoryController()
	if err != nil {
		log.Fatalf("Failed to initialize category controller: %v", err)
	}

	// 分类路由组
	categoryGroup := r.Group("/category")
	{
		categoryGroup.GET("/get", categoryController.GetCategory)     // 获取单个分类
		categoryGroup.GET("/list", categoryController.ListCategories) // 获取分类列表
		categoryGroup.POST("/create", categoryController.Create)      // 创建分类
		categoryGroup.POST("/update", categoryController.Update)      // 更新分类
		categoryGroup.POST("/delete", categoryController.Delete)      // 删除分类
	}
}
