// Package dto 提供 RBAC 权限管理相关的数据传输对象
package dto

// AddPolicyRequest 添加权限策略请求
type AddPolicyRequest struct {
	Role     string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// RemovePolicyRequest 删除权限策略请求
type RemovePolicyRequest struct {
	Role     string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// AddRoleForUserRequest 为用户分配角色请求
type AddRoleForUserRequest struct {
	User string `json:"user" validate:"required,min=1,max=100"` // 用户标识
	Role string `json:"role" validate:"required,min=1,max=100"` // 角色名称
}

// RemoveRoleForUserRequest 移除用户角色请求
type RemoveRoleForUserRequest struct {
	User string `json:"user" validate:"required,min=1,max=100"` // 用户标识
	Role string `json:"role" validate:"required,min=1,max=100"` // 角色名称
}

// GetRolesForUserRequest 获取用户角色请求
type GetRolesForUserRequest struct {
	User string `query:"user" validate:"required,min=1,max=100"` // 用户标识
}

// EnforceRequest 权限检查请求
type EnforceRequest struct {
	User     string `json:"user" validate:"required,min=1,max=100"`                 // 用户标识
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}
