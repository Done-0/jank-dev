// Package middleware provides common middleware for Hertz application.
package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/Done-0/jank/internal/middleware/cors"
	"github.com/Done-0/jank/internal/middleware/logger"
)

// New 初始化并注册所有中间件
// 参数
//
// h *server.Hertz: Hertz 服务器实例
func New(h *server.Hertz) {
	// 访问日志中间件
	h.Use(logger.New())

	// CORS 中间件
	h.Use(cors.New())
}
