package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// RBACService 权限管理服务接口
type RBACService interface {
	AddPolicy(ctx *app.RequestContext, req *dto.AddPolicyRequest) (*vo.PolicyResponse, error)                       // 添加策略
	RemovePolicy(ctx *app.RequestContext, req *dto.RemovePolicyRequest) (bool, error)                               // 删除策略
	AddRoleForUser(ctx *app.RequestContext, req *dto.AddRoleForUserRequest) (*vo.UserRoleResponse, error)           // 为用户添加角色
	RemoveRoleForUser(ctx *app.RequestContext, req *dto.RemoveRoleForUserRequest) (bool, error)                     // 删除用户角色
	AddRoleInheritance(ctx *app.RequestContext, req *dto.AddRoleInheritanceRequest) (*vo.RoleResponse, error)       // 添加角色继承关系
	RemoveRoleInheritance(ctx *app.RequestContext, req *dto.RemoveRoleInheritanceRequest) (bool, error)             // 删除角色继承关系
	GetRolesForUser(ctx *app.RequestContext, req *dto.GetRolesForUserRequest) (*vo.UserRoleResponse, error)         // 获取用户角色
	GetPoliciesForRole(ctx *app.RequestContext, req *dto.GetPoliciesForRoleRequest) (*vo.RolePolicyResponse, error) // 获取角色策略
	GetAllRoles(ctx *app.RequestContext, req *dto.GetAllRolesRequest) (*vo.RoleListResponse, error)                 // 获取所有角色
	GetAllPolicies(ctx *app.RequestContext, req *dto.GetAllPoliciesRequest) (*vo.PolicyListResponse, error)         // 获取所有策略
}
