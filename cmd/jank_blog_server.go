// Package cmd 提供应用程序的启动和运行入口
// 创建者：Done-0
// 创建时间：2025-08-05
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/db"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/logger"
	"github.com/Done-0/jank/internal/middleware"
	"github.com/Done-0/jank/internal/plugin"
	"github.com/Done-0/jank/internal/redis"
	"github.com/Done-0/jank/pkg/router"
)

// Start 启动服务
func Start() {
	// 初始化配置
	if err := configs.New(configs.DefaultConfigPath); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	config, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	// 初始化日志
	logger.New(config)

	// 初始化数据库
	db.New(config)

	// 初始化 Redis
	redis.New(config)

	// 初始化插件系统
	plugin.New(config)

	// 创建 Hertz 服务器实例
	addr := fmt.Sprintf("%s:%s", config.AppConfig.AppHost, config.AppConfig.AppPort)
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithExitWaitTime(10*time.Second),
	)

	// 注册中间件
	middleware.New(h)

	// 注册路由
	router.New(h)

	// 启动信息
	global.SysLog.Infof("⇨ Hertz server starting on %s", addr)

	// 启动服务器
	h.Spin()

	// 服务器关闭后清理插件
	global.SysLog.Info("Shutting down plugin system...")
	plugin.GlobalPluginManager.Shutdown()
	global.SysLog.Info("Plugin system shutdown completed")
}
