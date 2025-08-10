// Package errno 用户模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-05
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 用户模块错误码: 20000 ~ 29999
const (
	ErrUserNotFound      = 20001 // 用户不存在
	ErrUserInvalidPasswd = 20002 // 密码错误
	ErrUserTokenExpired  = 20003 // 令牌过期
	ErrUserBanned        = 20004 // 用户被禁用
	ErrUserExists        = 20005 // 用户已存在
	ErrUserInvalidEmail  = 20006 // 邮箱格式无效
	ErrUserWeakPassword  = 20007 // 密码强度过低
)

func init() {
	code.Register(ErrUserNotFound, "user not found: {id}")
	code.Register(ErrUserInvalidPasswd, "invalid password for user: {username}")
	code.Register(ErrUserTokenExpired, "token expired: {token}")
	code.Register(ErrUserBanned, "user banned: {username} reason: {reason}")
	code.Register(ErrUserExists, "user already exists: {username}")
	code.Register(ErrUserInvalidEmail, "invalid email format: {email}")
	code.Register(ErrUserWeakPassword, "password too weak: {msg}")
}
