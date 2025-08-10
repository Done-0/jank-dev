// Package mapper 提供RBAC权限管理相关的数据访问接口
// 创建者：Done-0
// 创建时间：2025-08-08
package mapper

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/rbac"
	"github.com/Done-0/jank/internal/model/user"
)

// RBACMapper RBAC权限管理数据访问接口
type RBACMapper interface {
	GetAllPolicies(ctx *app.RequestContext, pageNo, pageSize int64) ([]*rbac.Policy, int64, error) // GetAllPolicies 获取所有策略
	GetPoliciesForRole(ctx *app.RequestContext, role string) ([]*rbac.Policy, error)               // GetPoliciesForRole 获取角色的所有策略
	GetAllRoles(ctx *app.RequestContext, pageNo, pageSize int64) ([]string, int64, error)          // GetAllRoles 获取所有角色
	GetRoleInheritances(ctx *app.RequestContext, role string) ([]string, error)                    // GetRoleInheritances 获取角色继承关系
	GetUserByID(ctx *app.RequestContext, id int64) (*user.User, error)                             // GetUserByID 根据ID获取用户
	UpdateUserRole(ctx *app.RequestContext, userID int64, role string) error                       // UpdateUserRole 更新用户角色
}
