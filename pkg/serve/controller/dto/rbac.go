// Package dto 提供 RBAC 权限管理相关的数据传输对象
// 创建者：Done-0
// 创建时间：2025-08-08
package dto

// PolicyRequest 策略请求参数基础结构
type PolicyRequest struct {
	Role   string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Path   string `json:"path" validate:"required,min=1,max=255"`                 // 资源路径
	Method string `json:"method" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// AddPolicyRequest 添加策略请求参数
type AddPolicyRequest struct {
	PolicyRequest
}

// RemovePolicyRequest 删除策略请求参数
type RemovePolicyRequest struct {
	PolicyRequest
}

// AddRoleForUserRequest 为用户添加角色请求参数
type AddRoleForUserRequest struct {
	UserID int64  `json:"user_id,string" validate:"required,min=1"` // 用户ID
	Role   string `json:"role" validate:"required,min=1,max=100"`   // 角色名称
}

// RemoveRoleForUserRequest 删除用户角色请求参数
type RemoveRoleForUserRequest struct {
	UserID int64  `json:"user_id,string" validate:"required,min=1"` // 用户ID
	Role   string `json:"role" validate:"required,min=1,max=100"`   // 角色名称
}

// AddRoleInheritanceRequest 添加角色继承关系请求参数
type AddRoleInheritanceRequest struct {
	Role       string `json:"role" validate:"required,min=1,max=100"`        // 角色名称
	ParentRole string `json:"parent_role" validate:"required,min=1,max=100"` // 父角色名称
}

// RemoveRoleInheritanceRequest 删除角色继承关系请求参数
type RemoveRoleInheritanceRequest struct {
	Role       string `json:"role" validate:"required,min=1,max=100"`        // 角色名称
	ParentRole string `json:"parent_role" validate:"required,min=1,max=100"` // 父角色名称
}

// GetRolesForUserRequest 获取用户角色请求参数
type GetRolesForUserRequest struct {
	UserID int64 `query:"user_id,string" validate:"required,min=1"` // 用户ID
}

// GetPoliciesForRoleRequest 获取角色权限请求参数
type GetPoliciesForRoleRequest struct {
	Role string `query:"role" validate:"required,min=1,max=100"` // 角色名称
}

// GetAllRolesRequest 获取所有角色请求参数
type GetAllRolesRequest struct {
	PageNo   int64 `query:"page_no" validate:"omitempty,min=1"`           // 页码
	PageSize int64 `query:"page_size" validate:"omitempty,min=1,max=100"` // 每页数量
}

// GetAllPoliciesRequest 获取所有策略请求参数
type GetAllPoliciesRequest struct {
	PageNo   int64 `query:"page_no" validate:"omitempty,min=1"`           // 页码
	PageSize int64 `query:"page_size" validate:"omitempty,min=1,max=100"` // 每页数量
}
