// Package casbin 提供Casbin权限系统初始化和管理功能
// 创建者：Done-0
// 创建时间：2025-08-05
package casbin

import (
	"github.com/casbin/casbin/v2"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"

	gormAdapter "github.com/casbin/gorm-adapter/v3"
	hertzCasbin "github.com/hertz-contrib/casbin"
)

// New 初始化 Casbin 权限管理系统
// 参数：
//
//	config: 配置信息
func New(config *configs.Config) {
	var enforcer *casbin.Enforcer
	var err error

	// 根据适配器类型创建 Enforcer
	if config.CasbinConfig.DBAdapter {
		// 使用数据库适配器
		adapter, err := gormAdapter.NewAdapterByDB(global.DB)
		if err != nil {
			global.SysLog.Errorf("failed to create Casbin database adapter: %v", err)
			return
		}

		// 创建 Enforcer
		enforcer, err = casbin.NewEnforcer(config.CasbinConfig.ModelPath, adapter)
		if err != nil {
			global.SysLog.Errorf("failed to create Casbin Enforcer: %v", err)
			return
		}

		// 如果数据库中没有策略，从文件同步到数据库
		if policies, _ := enforcer.GetPolicy(); len(policies) == 0 {
			if fileEnforcer, err := casbin.NewEnforcer(config.CasbinConfig.ModelPath, config.CasbinConfig.PolicyPath); err == nil {
				// 清除现有策略
				enforcer.ClearPolicy()

				// 同步普通策略
				if policies, err := fileEnforcer.GetPolicy(); err == nil {
					for _, policy := range policies {
						if len(policy) >= 3 {
							params := make([]interface{}, len(policy))
							for i, v := range policy {
								params[i] = v
							}
							enforcer.AddPolicy(params...)
						}
					}
				}

				// 同步角色继承关系
				if groupPolicies, err := fileEnforcer.GetGroupingPolicy(); err == nil {
					for _, groupPolicy := range groupPolicies {
						if len(groupPolicy) >= 2 {
							params := make([]interface{}, len(groupPolicy))
							for i, v := range groupPolicy {
								params[i] = v
							}
							enforcer.AddGroupingPolicy(params...)
						}
					}
				}

				enforcer.SavePolicy()
				global.SysLog.Infof("Casbin policy synced from file to database")
			}
		}
	} else {
		// 使用文件适配器
		enforcer, err = casbin.NewEnforcer(config.CasbinConfig.ModelPath, config.CasbinConfig.PolicyPath)
		if err != nil {
			global.SysLog.Errorf("创建 Casbin Enforcer 失败: %v", err)
			return
		}
	}

	// 启用自动保存
	enforcer.EnableAutoSave(true)
	global.Enforcer = enforcer
	global.SysLog.Infof("Casbin enforcer initialized successfully")
}

// GetMiddleware 获取 Hertz Casbin 中间件
// 参数：
//
//	modelPath: 模型文件路径
//	lookupHandler: 查找处理器
//
// 返回值：
//
//	*hertzCasbin.Middleware: Casbin 中间件
//	error: 错误信息
func GetMiddleware(modelPath string, lookupHandler hertzCasbin.LookupHandler) (*hertzCasbin.Middleware, error) {
	return hertzCasbin.NewCasbinMiddleware(modelPath, global.Enforcer, lookupHandler)
}
