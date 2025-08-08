// Package impl 提供RBAC权限管理相关的服务层实现
// 创建者：Done-0
// 创建时间：2025-08-08
package impl

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

// RBACServiceImpl RBAC权限管理服务实现
type RBACServiceImpl struct {
	rbacMapper mapper.RBACMapper
}

// NewRBACService 创建RBAC权限管理服务实例
// 参数：
//
//	mapper: RBAC权限管理数据访问接口
//
// 返回值：
//
//	service.RBACService: RBAC权限管理服务接口
func NewRBACService(mapper mapper.RBACMapper) service.RBACService {
	return &RBACServiceImpl{rbacMapper: mapper}
}

// AddPolicy 添加策略
func (s *RBACServiceImpl) AddPolicy(ctx *app.RequestContext, req *dto.AddPolicyRequest) (*vo.PolicyResponse, error) {
	ok, err := global.Enforcer.AddPolicy(req.Role, req.Path, req.Method)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to add policy, role[%s], path[%s], method[%s]: %v", req.Role, req.Path, req.Method, err)
		return nil, fmt.Errorf("failed to add policy: %w", err)
	}

	if !ok {
		logger.BizLogger(ctx).Warnf("policy already exists, role[%s], path[%s], method[%s]", req.Role, req.Path, req.Method)
		return nil, fmt.Errorf("policy already exists")
	}

	if err = global.Enforcer.SavePolicy(); err != nil {
		logger.BizLogger(ctx).Errorf("failed to save policy: %v", err)
		return nil, fmt.Errorf("failed to save policy: %w", err)
	}

	return &vo.PolicyResponse{
		Role:   req.Role,
		Path:   req.Path,
		Method: req.Method,
	}, nil
}

// RemovePolicy 删除策略
func (s *RBACServiceImpl) RemovePolicy(ctx *app.RequestContext, req *dto.RemovePolicyRequest) (bool, error) {
	ok, err := global.Enforcer.RemovePolicy(req.Role, req.Path, req.Method)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to remove policy, role[%s], path[%s], method[%s]: %v", req.Role, req.Path, req.Method, err)
		return false, fmt.Errorf("failed to remove policy: %w", err)
	}

	if !ok {
		logger.BizLogger(ctx).Warnf("policy not found, role[%s], path[%s], method[%s]", req.Role, req.Path, req.Method)
		return false, nil
	}

	if err = global.Enforcer.SavePolicy(); err != nil {
		logger.BizLogger(ctx).Errorf("failed to save policy: %v", err)
		return false, fmt.Errorf("failed to save policy: %w", err)
	}

	return true, nil
}

// AddRoleForUser 为用户添加角色
func (s *RBACServiceImpl) AddRoleForUser(ctx *app.RequestContext, req *dto.AddRoleForUserRequest) (*vo.UserRoleResponse, error) {
	_, err := s.rbacMapper.GetUserByID(ctx, req.UserID)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get user info, userID[%d]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	if err = s.rbacMapper.UpdateUserRole(ctx, req.UserID, req.Role); err != nil {
		logger.BizLogger(ctx).Errorf("failed to update user role, userID[%d], role[%s]: %v", req.UserID, req.Role, err)
		return nil, fmt.Errorf("failed to update user role: %w", err)
	}

	inherited, err := s.rbacMapper.GetRoleInheritances(ctx, req.Role)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get role inheritances, role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to get role inheritances: %w", err)
	}

	roles := []string{req.Role}
	roles = append(roles, inherited...)

	return &vo.UserRoleResponse{
		UserID: req.UserID,
		Roles:  roles,
	}, nil
}

// RemoveRoleForUser 删除用户角色
func (s *RBACServiceImpl) RemoveRoleForUser(ctx *app.RequestContext, req *dto.RemoveRoleForUserRequest) (bool, error) {
	u, err := s.rbacMapper.GetUserByID(ctx, req.UserID)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get user info, userID[%d]: %v", req.UserID, err)
		return false, fmt.Errorf("failed to get user info: %w", err)
	}

	if u.Role != req.Role {
		logger.BizLogger(ctx).Warnf("user current role is not the role to be removed, userID[%d], current role[%s], target role[%s]", req.UserID, u.Role, req.Role)
		return false, nil
	}

	if err = s.rbacMapper.UpdateUserRole(ctx, req.UserID, "user"); err != nil {
		logger.BizLogger(ctx).Errorf("failed to update user role, userID[%d], role[user]: %v", req.UserID, err)
		return false, fmt.Errorf("failed to update user role: %w", err)
	}

	return true, nil
}

// AddRoleInheritance 添加角色继承关系
func (s *RBACServiceImpl) AddRoleInheritance(ctx *app.RequestContext, req *dto.AddRoleInheritanceRequest) (*vo.RoleResponse, error) {
	ok, err := global.Enforcer.AddRoleForUser(req.Role, req.ParentRole)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to add role inheritance, role[%s], parent role[%s]: %v", req.Role, req.ParentRole, err)
		return nil, fmt.Errorf("failed to add role inheritance: %w", err)
	}

	if !ok {
		logger.BizLogger(ctx).Warnf("role inheritance already exists, role[%s], parent role[%s]", req.Role, req.ParentRole)
		return nil, fmt.Errorf("role inheritance already exists")
	}

	if err = global.Enforcer.SavePolicy(); err != nil {
		logger.BizLogger(ctx).Errorf("failed to save policy: %v", err)
		return nil, fmt.Errorf("failed to save policy: %w", err)
	}

	inherited, err := s.rbacMapper.GetRoleInheritances(ctx, req.Role)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get role inheritances, role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to get role inheritances: %w", err)
	}

	return &vo.RoleResponse{
		Role:           req.Role,
		InheritedRoles: inherited,
	}, nil
}

// RemoveRoleInheritance 删除角色继承关系
func (s *RBACServiceImpl) RemoveRoleInheritance(ctx *app.RequestContext, req *dto.RemoveRoleInheritanceRequest) (bool, error) {
	ok, err := global.Enforcer.DeleteRoleForUser(req.Role, req.ParentRole)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to remove role inheritance, role[%s], parent role[%s]: %v", req.Role, req.ParentRole, err)
		return false, fmt.Errorf("failed to remove role inheritance: %w", err)
	}

	if !ok {
		logger.BizLogger(ctx).Warnf("role inheritance does not exist, role[%s], parent role[%s]", req.Role, req.ParentRole)
		return false, nil
	}

	if err = global.Enforcer.SavePolicy(); err != nil {
		logger.BizLogger(ctx).Errorf("failed to save policy: %v", err)
		return false, fmt.Errorf("failed to save policy: %w", err)
	}

	return true, nil
}

// GetRolesForUser 获取用户角色
func (s *RBACServiceImpl) GetRolesForUser(ctx *app.RequestContext, req *dto.GetRolesForUserRequest) (*vo.UserRoleResponse, error) {
	u, err := s.rbacMapper.GetUserByID(ctx, req.UserID)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get user info, userID[%d]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	roles := []string{u.Role}

	inherited, err := s.rbacMapper.GetRoleInheritances(ctx, u.Role)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get role inheritances, role[%s]: %v", u.Role, err)
		return nil, fmt.Errorf("failed to get role inheritances: %w", err)
	}

	roles = append(roles, inherited...)

	return &vo.UserRoleResponse{
		UserID: req.UserID,
		Roles:  roles,
	}, nil
}

// GetPoliciesForRole 获取角色策略
func (s *RBACServiceImpl) GetPoliciesForRole(ctx *app.RequestContext, req *dto.GetPoliciesForRoleRequest) (*vo.RolePolicyResponse, error) {
	policies, err := s.rbacMapper.GetPoliciesForRole(ctx, req.Role)
	if err != nil {
		logger.BizLogger(ctx).Errorf("failed to get policies for role, role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to get policies for role: %w", err)
	}

	list := make([]vo.PolicyResponse, 0, len(policies))
	for _, policy := range policies {
		list = append(list, vo.PolicyResponse{
			Role:   policy.V0,
			Path:   policy.V1,
			Method: policy.V2,
		})
	}

	return &vo.RolePolicyResponse{
		Role:     req.Role,
		Policies: list,
	}, nil
}

// GetAllRoles 获取所有角色
func (s *RBACServiceImpl) GetAllRoles(ctx *app.RequestContext, req *dto.GetAllRolesRequest) (*vo.RoleListResponse, error) {
	roles, total, err := s.rbacMapper.GetAllRoles(ctx, req.PageNo, req.PageSize)
	if err != nil {
		global.SysLog.Errorf("获取所有角色失败: %v", err)
		return nil, fmt.Errorf("获取所有角色失败: %w", err)
	}

	list := make([]vo.RoleResponse, 0, len(roles))
	for _, role := range roles {
		inherited, err := s.rbacMapper.GetRoleInheritances(ctx, role)
		if err != nil {
			global.SysLog.Errorf("获取角色继承关系失败，角色[%s]: %v", role, err)
			return nil, fmt.Errorf("获取角色继承关系失败: %w", err)
		}

		list = append(list, vo.RoleResponse{
			Role:           role,
			InheritedRoles: inherited,
		})
	}

	return &vo.RoleListResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}

// GetAllPolicies 获取所有策略
func (s *RBACServiceImpl) GetAllPolicies(ctx *app.RequestContext, req *dto.GetAllPoliciesRequest) (*vo.PolicyListResponse, error) {
	policies, total, err := s.rbacMapper.GetAllPolicies(ctx, req.PageNo, req.PageSize)
	if err != nil {
		global.SysLog.Errorf("获取所有策略失败: %v", err)
		return nil, fmt.Errorf("获取所有策略失败: %w", err)
	}

	list := make([]vo.PolicyResponse, 0, len(policies))
	for _, policy := range policies {
		list = append(list, vo.PolicyResponse{
			Role:   policy.V0,
			Path:   policy.V1,
			Method: policy.V2,
		})
	}

	return &vo.PolicyListResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}
