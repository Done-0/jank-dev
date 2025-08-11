// Package consts 提供验证码相关常量定义
// 创建者：Done-0
// 创建时间：2025-08-10
package consts

import "time"

const (
	// 验证码缓存键前缀
	EmailVerificationCodeCacheKeyPrefix = "EMAIL:VERIFICATION:CODE"     // 邮箱验证码缓存键前缀
	ImgVerificationCodeCacheKeyPrefix   = "IMG:VERIFICATION:CODE:CACHE" // 图形验证码缓存键前缀

	// 验证码缓存过期时间
	EmailVerificationCodeCacheExpiration = 3 * time.Minute // 邮箱验证码缓存过期时间（3分钟）
	ImgVerificationCodeCacheExpiration   = 5 * time.Minute // 图形验证码缓存过期时间（5分钟）

	// 验证码长度配置
	EmailVerificationCodeLength = 6 // 邮箱验证码长度（6位数字）
	ImgVerificationCodeLength   = 4 // 图形验证码长度（4位字符）
)
