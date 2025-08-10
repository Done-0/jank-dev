// Package rbac 提供 RBAC 权限管理相关的视图对象，包含策略、角色等权限控制功能定义
// 创建者：Done-0
// 创建时间：2025-08-08
package vo

// PolicyResponse 角色权限策略响应
type PolicyResponse struct {
	Role   string `json:"role"`   // 角色名称
	Path   string `json:"path"`   // 资源路径
	Method string `json:"method"` // 操作方法
}

// RoleResponse 角色响应
type RoleResponse struct {
	Role           string   `json:"role"`            // 角色名称
	InheritedRoles []string `json:"inherited_roles"` // 继承的角色列表
}

// UserRoleResponse 用户角色响应
type UserRoleResponse struct {
	UserID int64    `json:"user_id"` // 用户 ID
	Roles  []string `json:"roles"`   // 角色列表
}

// RolePolicyResponse 角色权限策略响应
type RolePolicyResponse struct {
	Role     string           `json:"role"`     // 角色名称
	Policies []PolicyResponse `json:"policies"` // 策略列表
}

// PolicyListResponse 策略列表响应
type PolicyListResponse struct {
	Total    int64            `json:"total"`     // 总条数
	PageNo   int64            `json:"page_no"`   // 当前页
	PageSize int64            `json:"page_size"` // 当前分页记录数
	List     []PolicyResponse `json:"list"`      // 分页内容
}

// RoleListResponse 角色列表响应
type RoleListResponse struct {
	Total    int64          `json:"total"`     // 总条数
	PageNo   int64          `json:"page_no"`   // 当前页
	PageSize int64          `json:"page_size"` // 当前分页记录数
	List     []RoleResponse `json:"list"`      // 分页内容
}
