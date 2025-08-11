package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// RBACService 权限管理服务接口
type RBACService interface {
	// 策略管理
	AddPolicy(c *app.RequestContext, req *dto.AddPolicyRequest) (*vo.PolicyResponse, error) // 添加权限策略
	RemovePolicy(c *app.RequestContext, req *dto.RemovePolicyRequest) (bool, error)         // 删除权限策略
	GetAllPolicies(c *app.RequestContext) (*vo.PolicyListResponse, error)                   // 获取所有策略

	// 角色分配
	AddRoleForUser(c *app.RequestContext, req *dto.AddRoleForUserRequest) (*vo.UserRolesResponse, error)   // 为用户分配角色
	RemoveRoleForUser(c *app.RequestContext, req *dto.RemoveRoleForUserRequest) (bool, error)              // 移除用户角色
	GetRolesForUser(c *app.RequestContext, req *dto.GetRolesForUserRequest) (*vo.UserRolesResponse, error) // 获取用户角色

	// 权限检查
	Enforce(c *app.RequestContext, req *dto.EnforceRequest) (*vo.EnforceResponse, error) // 权限检查
}
