// Package consts 提供缓存相关常量定义
// 创建者：Done-0
// 创建时间：2025-08-10
package consts

const (
	// Redis 缓存键前缀
	UserCacheKeyPrefix        = "USER:CACHE"         // 用户 access token 缓存键前缀
	UserRefreshCacheKeyPrefix = "USER:REFRESH:CACHE" // 用户 refresh token 缓存键前缀
)
