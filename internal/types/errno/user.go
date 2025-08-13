// Package errno 用户模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-05
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 用户模块错误码: 50000 ~ 59999
const (
	ErrUserRegisterFailed      = 20001 // 注册失败
	ErrUserLoginFailed         = 20002 // 登录失败
	ErrUserLogoutFailed        = 20003 // 登出失败
	ErrUserGetProfileFailed    = 20004 // 获取用户资料失败
	ErrUserResetPasswordFailed = 20005 // 重置密码失败
	ErrUserListFailed          = 20006 // 获取用户列表失败
	ErrUserRefreshTokenFailed  = 20007 // 刷新 token 失败
)

func init() {
	code.Register(ErrUserRegisterFailed, "user registration failed: {email}")
	code.Register(ErrUserLoginFailed, "user login failed: {email}")
	code.Register(ErrUserLogoutFailed, "user logout failed: {msg}")
	code.Register(ErrUserGetProfileFailed, "get user profile failed: {msg}")
	code.Register(ErrUserResetPasswordFailed, "reset password failed: {msg}")
	code.Register(ErrUserListFailed, "list users failed: {msg}")
	code.Register(ErrUserRefreshTokenFailed, "refresh token failed: {msg}")
}
