// Package impl 提供RBAC权限相关的数据访问实现
// 创建者：Done-0
// 创建时间：2025-08-11
package impl

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model/rbac"
	"github.com/Done-0/jank/internal/utils/db"
	"github.com/Done-0/jank/pkg/serve/mapper"
)

// RBACMapperImpl RBAC数据访问实现
type RBACMapperImpl struct{}

// NewRBACMapper 创建RBAC数据访问实例
func NewRBACMapper() mapper.RBACMapper {
	return &RBACMapperImpl{}
}

// AddPolicy 添加权限策略
func (m *RBACMapperImpl) AddPolicy(c *app.RequestContext, role, resource, action string) (*rbac.Policy, error) {
	exists, err := m.PolicyExists(c, role, resource, action)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, nil
	}

	policy := &rbac.Policy{
		Ptype: "p",
		V0:    role,
		V1:    resource,
		V2:    action,
	}

	err = db.GetDBFromContext(c).Create(policy).Error
	if err != nil {
		return nil, err
	}

	global.Enforcer.LoadPolicy()
	return policy, nil
}

// RemovePolicy 删除权限策略
func (m *RBACMapperImpl) RemovePolicy(c *app.RequestContext, role, resource, action string) (bool, error) {
	result := db.GetDBFromContext(c).Where("ptype = ? AND v0 = ? AND v1 = ? AND v2 = ?", "p", role, resource, action).Delete(&rbac.Policy{})
	if result.Error != nil {
		return false, result.Error
	}
	global.Enforcer.LoadPolicy()
	return result.RowsAffected > 0, nil
}

// GetAllPolicies 获取所有策略
func (m *RBACMapperImpl) GetAllPolicies(c *app.RequestContext) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).Find(&policies).Error
	return policies, err
}

// PolicyExists 检查策略是否存在
func (m *RBACMapperImpl) PolicyExists(c *app.RequestContext, role, resource, action string) (bool, error) {
	var count int64
	err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ? AND v2 = ?", "p", role, resource, action).Count(&count).Error
	return count > 0, err
}

// AddRoleForUser 为用户分配角色
func (m *RBACMapperImpl) AddRoleForUser(c *app.RequestContext, user, role string) (*rbac.Policy, error) {
	has, err := m.UserHasRole(c, user, role)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, nil
	}

	policy := &rbac.Policy{
		Ptype: "g",
		V0:    user,
		V1:    role,
	}

	err = db.GetDBFromContext(c).Create(policy).Error
	if err != nil {
		return nil, err
	}

	global.Enforcer.LoadPolicy()
	return policy, nil
}

// RemoveRoleForUser 移除用户角色
func (m *RBACMapperImpl) RemoveRoleForUser(c *app.RequestContext, user, role string) (bool, error) {
	result := db.GetDBFromContext(c).Where("ptype = ? AND v0 = ? AND v1 = ?", "g", user, role).Delete(&rbac.Policy{})
	if result.Error != nil {
		return false, result.Error
	}
	global.Enforcer.LoadPolicy()
	return result.RowsAffected > 0, nil
}

// GetRolesForUser 获取用户角色
func (m *RBACMapperImpl) GetRolesForUser(c *app.RequestContext, user string) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).Where("ptype = ? AND v0 = ?", "g", user).Find(&policies).Error
	return policies, err
}

// UserHasRole 检查用户是否有指定角色
func (m *RBACMapperImpl) UserHasRole(c *app.RequestContext, user, role string) (bool, error) {
	var count int64
	err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ?", "g", user, role).Count(&count).Error
	return count > 0, err
}

// Enforce 权限检查
func (m *RBACMapperImpl) Enforce(c *app.RequestContext, user, resource, action string) (bool, error) {
	return global.Enforcer.Enforce(user, resource, action)
}
