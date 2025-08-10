// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"log"

	"github.com/Done-0/jank/pkg/wire"
	"github.com/cloudwego/hertz/pkg/route"
)

// RegisterRBACRoutes 注册RBAC相关路由
func RegisterRBACRoutes(r *route.RouterGroup) {
	rbacController, err := wire.NewRBACController()
	if err != nil {
		log.Fatalf("Failed to initialize RBAC controller: %v", err)
	}

	// RBAC 路由组
	rbacGroup := r.Group("/rbac")
	{
		// 策略管理
		rbacGroup.POST("/addPolicy", rbacController.AddPolicy)          // 添加策略
		rbacGroup.POST("/removePolicy", rbacController.RemovePolicy)    // 删除策略
		rbacGroup.GET("/getAllPolicies", rbacController.GetAllPolicies) // 获取所有策略

		// 角色管理
		rbacGroup.GET("/getAllRoles", rbacController.GetAllRoles)                      // 获取所有角色
		rbacGroup.GET("/getPoliciesForRole", rbacController.GetPoliciesForRole)        // 获取角色策略
		rbacGroup.POST("/addRoleInheritance", rbacController.AddRoleInheritance)       // 添加角色继承
		rbacGroup.POST("/removeRoleInheritance", rbacController.RemoveRoleInheritance) // 删除角色继承

		// 用户角色管理
		rbacGroup.POST("/addRoleForUser", rbacController.AddRoleForUser)       // 添加用户角色
		rbacGroup.POST("/removeRoleForUser", rbacController.RemoveRoleForUser) // 删除用户角色
		rbacGroup.GET("/getRolesForUser", rbacController.GetRolesForUser)      // 获取用户角色
	}
}
