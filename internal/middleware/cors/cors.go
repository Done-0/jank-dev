// Package cors 提供CORS跨域中间件配置
// 创建者：Done-0
// 创建时间：2025-08-05
package cors

import (
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"

	"github.com/Done-0/jank/configs"
)

// New 创建 CORS 中间件
// 返回值：
// app.HandlerFunc: CORS中间件
func New() app.HandlerFunc {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	return cors.New(cors.Config{
		AllowOrigins:     cfgs.AppConfig.CORSConfig.AllowOrigins,
		AllowMethods:     cfgs.AppConfig.CORSConfig.AllowMethods,
		AllowHeaders:     cfgs.AppConfig.CORSConfig.AllowHeaders,
		ExposeHeaders:    cfgs.AppConfig.CORSConfig.ExposeHeaders,
		AllowCredentials: cfgs.AppConfig.CORSConfig.AllowCredentials,
		MaxAge:           time.Duration(cfgs.AppConfig.CORSConfig.MaxAge) * time.Hour,
	})
}
