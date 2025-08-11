package impl

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

// RBACServiceImpl RBAC服务实现
type RBACServiceImpl struct {
	rbacMapper mapper.RBACMapper
}

// NewRBACService 创建RBAC服务实例
func NewRBACService(rbacMapper mapper.RBACMapper) service.RBACService {
	return &RBACServiceImpl{
		rbacMapper: rbacMapper,
	}
}

// AddPolicy 添加权限策略
func (s *RBACServiceImpl) AddPolicy(c *app.RequestContext, req *dto.AddPolicyRequest) (*vo.PolicyResponse, error) {
	policy, err := s.rbacMapper.AddPolicy(c, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to add policy: role[%s], resource[%s], action[%s]: %v", req.Role, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to add policy: %w", err)
	}
	if policy == nil {
		logger.BizLogger(c).Warnf("policy already exists: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
		return nil, fmt.Errorf("policy already exists")
	}
	return &vo.PolicyResponse{
		Role:     req.Role,
		Resource: req.Resource,
		Action:   req.Action,
	}, nil
}

// RemovePolicy 删除权限策略
func (s *RBACServiceImpl) RemovePolicy(c *app.RequestContext, req *dto.RemovePolicyRequest) (bool, error) {
	ok, err := s.rbacMapper.RemovePolicy(c, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to remove policy: role[%s], resource[%s], action[%s]: %v", req.Role, req.Resource, req.Action, err)
		return false, fmt.Errorf("failed to remove policy: %w", err)
	}
	if !ok {
		logger.BizLogger(c).Warnf("policy not found: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
	}
	return ok, nil
}

// GetAllPolicies 获取所有策略
func (s *RBACServiceImpl) GetAllPolicies(c *app.RequestContext) (*vo.PolicyListResponse, error) {
	policies, err := s.rbacMapper.GetAllPolicies(c)
	if err != nil {
		logger.BizLogger(c).Errorf("Failed to get policies: %v", err)
		return nil, fmt.Errorf("failed to get policies: %w", err)
	}

	list := make([]vo.PolicyResponse, 0, len(policies))
	for _, policy := range policies {
		list = append(list, vo.PolicyResponse{
			Role:     policy.V0,
			Resource: policy.V1,
			Action:   policy.V2,
		})
	}
	return &vo.PolicyListResponse{
		List: list,
	}, nil
}

// AddRoleForUser 为用户分配角色
func (s *RBACServiceImpl) AddRoleForUser(c *app.RequestContext, req *dto.AddRoleForUserRequest) (*vo.UserRolesResponse, error) {
	policy, err := s.rbacMapper.AddRoleForUser(c, req.User, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to add role for user: user[%s], role[%s]: %v", req.User, req.Role, err)
		return nil, fmt.Errorf("failed to add role for user: %w", err)
	}
	if policy == nil {
		logger.BizLogger(c).Warnf("user already has role: user[%s], role[%s]", req.User, req.Role)
	}
	rolePolicies, err := s.rbacMapper.GetRolesForUser(c, req.User)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get roles for user: user[%s]: %v", req.User, err)
		return nil, fmt.Errorf("failed to get roles for user: %w", err)
	}

	roles := make([]string, 0, len(rolePolicies))
	for _, rolePolicy := range rolePolicies {
		roles = append(roles, rolePolicy.V1)
	}

	return &vo.UserRolesResponse{
		User:  req.User,
		Roles: roles,
	}, nil
}

// RemoveRoleForUser 移除用户角色
func (s *RBACServiceImpl) RemoveRoleForUser(c *app.RequestContext, req *dto.RemoveRoleForUserRequest) (bool, error) {
	ok, err := s.rbacMapper.RemoveRoleForUser(c, req.User, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to remove role for user: user[%s], role[%s]: %v", req.User, req.Role, err)
		return false, fmt.Errorf("failed to remove role for user: %w", err)
	}
	if !ok {
		logger.BizLogger(c).Warnf("user does not have role: user[%s], role[%s]", req.User, req.Role)
	}
	return ok, nil
}

// GetRolesForUser 获取用户角色
func (s *RBACServiceImpl) GetRolesForUser(c *app.RequestContext, req *dto.GetRolesForUserRequest) (*vo.UserRolesResponse, error) {
	rolePolicies, err := s.rbacMapper.GetRolesForUser(c, req.User)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get roles for user: user[%s]: %v", req.User, err)
		return nil, fmt.Errorf("failed to get roles for user: %w", err)
	}

	roles := make([]string, 0, len(rolePolicies))
	for _, rolePolicy := range rolePolicies {
		roles = append(roles, rolePolicy.V1)
	}

	return &vo.UserRolesResponse{
		User:  req.User,
		Roles: roles,
	}, nil
}

// Enforce 权限检查
func (s *RBACServiceImpl) Enforce(c *app.RequestContext, req *dto.EnforceRequest) (*vo.EnforceResponse, error) {
	allowed, err := s.rbacMapper.Enforce(c, req.User, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to enforce: user[%s], resource[%s], action[%s]: %v", req.User, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to enforce: %w", err)
	}
	reason := "denied"
	if allowed {
		reason = "allowed"
	}
	return &vo.EnforceResponse{
		Allowed: allowed,
		Reason:  reason,
	}, nil
}
