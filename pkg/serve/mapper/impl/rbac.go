// Package impl 提供RBAC权限管理相关的数据访问实现
// 创建者：Done-0
// 创建时间：2025-08-08
package impl

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model/rbac"
	"github.com/Done-0/jank/internal/model/user"
	"github.com/Done-0/jank/pkg/serve/mapper"
)

// RBACMapperImpl RBAC权限管理数据访问实现
type RBACMapperImpl struct{}

// NewRBACMapper 创建RBAC权限管理数据访问实例
// 返回值：
//
//	mapper.RBACMapper: RBAC权限管理数据访问接口
func NewRBACMapper() mapper.RBACMapper {
	return &RBACMapperImpl{}
}

// GetAllPolicies 获取所有策略
// 参数：
//
//	ctx: 上下文信息
//	pageNo: 页码
//	pageSize: 每页数量
//
// 返回值：
//
//	[]*rbac.Policy: 策略列表
//	int64: 总记录数
//	error: 错误信息
func (m *RBACMapperImpl) GetAllPolicies(ctx *app.RequestContext, pageNo, pageSize int64) ([]*rbac.Policy, int64, error) {
	var policies []*rbac.Policy
	var total int64

	db := global.DB
	query := db.Model(&rbac.Policy{}).Where("ptype = ?", "p")

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get total policies: %w", err)
	}

	// 查询记录 - 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.
		Order("id DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&policies).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get policy list: %w", err)
	}

	return policies, total, nil
}

// GetPoliciesForRole 获取角色的所有策略
// 参数：
//
//	ctx: 上下文信息
//	role: 角色名称
//
// 返回值：
//
//	[]*rbac.Policy: 策略列表
//	error: 错误信息
func (m *RBACMapperImpl) GetPoliciesForRole(ctx *app.RequestContext, role string) ([]*rbac.Policy, error) {
	db := global.DB

	var policies []*rbac.Policy
	if err := db.Model(&rbac.Policy{}).
		Where("ptype = ? AND v0 = ?", "p", role).
		Find(&policies).Error; err != nil {
		return nil, fmt.Errorf("failed to get role policies: %w", err)
	}

	return policies, nil
}

// GetAllRoles 获取所有角色
// 参数：
//
//	ctx: 上下文信息
//	pageNo: 页码
//	pageSize: 每页数量
//
// 返回值：
//
//	[]string: 角色列表
//	int64: 总记录数
//	error: 错误信息
func (m *RBACMapperImpl) GetAllRoles(ctx *app.RequestContext, pageNo, pageSize int64) ([]string, int64, error) {
	var policies []*rbac.Policy
	var total int64

	db := global.DB
	query := db.Model(&rbac.Policy{}).
		Select("DISTINCT v0").
		Where("ptype = ?", "p")

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get total roles: %w", err)
	}

	// 查询记录 - 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&policies).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get role list: %w", err)
	}

	// 提取角色名称
	roles := make([]string, 0, len(policies))
	for _, policy := range policies {
		roles = append(roles, policy.V0)
	}

	return roles, total, nil
}

// GetRoleInheritances 获取角色继承关系
// 参数：
//
//	ctx: 上下文信息
//	role: 角色名称
//
// 返回值：
//
//	[]string: 继承的角色列表
//	error: 错误信息
func (m *RBACMapperImpl) GetRoleInheritances(ctx *app.RequestContext, role string) ([]string, error) {
	db := global.DB

	var policies []*rbac.Policy
	if err := db.Model(&rbac.Policy{}).
		Where("ptype = ? AND v0 = ?", "g", role).
		Find(&policies).Error; err != nil {
		return nil, fmt.Errorf("failed to get role inheritances: %w", err)
	}

	// 提取继承的角色
	inheritedRoles := make([]string, 0, len(policies))
	for _, policy := range policies {
		inheritedRoles = append(inheritedRoles, policy.V1)
	}

	return inheritedRoles, nil
}

// GetUserByID 根据ID获取用户
// 参数：
//
//	ctx: 上下文信息
//	id: 用户ID
//
// 返回值：
//
//	*user.User: 用户信息
//	error: 错误信息
func (m *RBACMapperImpl) GetUserByID(ctx *app.RequestContext, id int64) (*user.User, error) {
	var u user.User
	db := global.DB

	err := db.First(&u, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	return &u, nil
}

// UpdateUserRole 更新用户角色
// 参数：
//
//	ctx: 上下文信息
//	userID: 用户ID
//	role: 角色名称
//
// 返回值：
//
//	error: 错误信息
func (m *RBACMapperImpl) UpdateUserRole(ctx *app.RequestContext, userID int64, role string) error {
	db := global.DB

	// 更新用户角色
	if err := db.Model(&user.User{}).
		Where("id = ?", userID).
		Update("role", role).Error; err != nil {
		return fmt.Errorf("failed to update user role: %w", err)
	}

	return nil
}
