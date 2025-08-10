// Package errno RBAC权限模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-05
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// RBAC模块错误码: 30000 ~ 39999
const (
	ErrPermissionDenied = 30001 // 权限被拒绝
	ErrRoleNotFound     = 30002 // 角色不存在
	ErrRoleExists       = 30003 // 角色已存在
	ErrPolicyNotFound   = 30004 // 策略不存在
	ErrPolicyExists     = 30005 // 策略已存在
	ErrInvalidResource  = 30006 // 无效的资源
	ErrInvalidAction    = 30007 // 无效的操作
	ErrUserRoleNotFound = 30008 // 用户角色不存在
)

func init() {
	code.Register(ErrPermissionDenied, "permission denied for {user} on {resource}: {action}")
	code.Register(ErrRoleNotFound, "role not found: {role}")
	code.Register(ErrRoleExists, "role already exists: {role}")
	code.Register(ErrPolicyNotFound, "policy not found: {policy}")
	code.Register(ErrPolicyExists, "policy already exists: {policy}")
	code.Register(ErrInvalidResource, "invalid resource: {resource}")
	code.Register(ErrInvalidAction, "invalid action: {action} on {resource}")
	code.Register(ErrUserRoleNotFound, "user role not found: {user} -> {role}")
}
