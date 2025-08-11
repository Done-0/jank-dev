// Package mapper 提供RBAC权限相关的数据访问接口
// 创建者：Done-0
// 创建时间：2025-08-11
package mapper

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/rbac"
)

// RBACMapper RBAC数据访问接口
type RBACMapper interface {
	// 策略管理
	AddPolicy(c *app.RequestContext, role, resource, action string) (*rbac.Policy, error) // 添加权限策略
	RemovePolicy(c *app.RequestContext, role, resource, action string) (bool, error)      // 删除权限策略
	GetAllPolicies(c *app.RequestContext) ([]*rbac.Policy, error)                         // 获取所有策略
	PolicyExists(c *app.RequestContext, role, resource, action string) (bool, error)      // 检查策略是否存在

	// 角色分配
	AddRoleForUser(c *app.RequestContext, user, role string) (*rbac.Policy, error) // 为用户分配角色
	RemoveRoleForUser(c *app.RequestContext, user, role string) (bool, error)      // 移除用户角色
	GetRolesForUser(c *app.RequestContext, user string) ([]*rbac.Policy, error)    // 获取用户角色
	UserHasRole(c *app.RequestContext, user, role string) (bool, error)            // 检查用户是否有指定角色

	// 权限检查
	Enforce(c *app.RequestContext, user, resource, action string) (bool, error) // 权限检查
}
