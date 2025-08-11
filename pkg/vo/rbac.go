// Package vo RBAC 相关值对象
// 创建者：Done-0
// 创建时间：2025-08-08
package vo

// PolicyResponse 权限策略响应 (p, role, resource, action)
type PolicyResponse struct {
	Role     string `json:"role"`     // 角色名称
	Resource string `json:"resource"` // 资源路径
	Action   string `json:"action"`   // 操作方法
}

// UserRolesResponse 用户角色响应
type UserRolesResponse struct {
	User  string   `json:"user"`  // 用户标识
	Roles []string `json:"roles"` // 角色列表
}

// EnforceResponse 权限检查响应
type EnforceResponse struct {
	Allowed bool   `json:"allowed"` // 是否允许
	Reason  string `json:"reason"`  // 原因说明
}

// PolicyListResponse 策略列表响应
type PolicyListResponse struct {
	Total    int64            `json:"total"`     // 总条数
	PageNo   int64            `json:"page_no"`   // 当前页
	PageSize int64            `json:"page_size"` // 当前分页记录数
	List     []PolicyResponse `json:"list"`      // 分页内容
}
