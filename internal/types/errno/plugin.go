// Package errno 插件模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-06
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 插件模块错误码: 40000 ~ 49999
const (
	ErrPluginNotFound           = 40001 // 插件不存在
	ErrPluginExists             = 40002 // 插件已存在
	ErrPluginNotRunning         = 40003 // 插件未运行
	ErrPluginStartFailed        = 40004 // 插件启动失败
	ErrPluginStopFailed         = 40005 // 插件停止失败
	ErrPluginInvalidConfig      = 40006 // 插件配置无效
	ErrPluginTimeout            = 40007 // 操作超时
	ErrPluginInvalidOpenapi3Doc = 40008 // OpenAPI 3.0 文档无效
	ErrPluginGRPCConnFailed     = 40009 // gRPC 连接失败
	ErrPluginInstallFailed      = 40010 // 插件安装失败
	ErrPluginUninstallFailed    = 40011 // 插件卸载失败
	ErrPluginRegisterFailed     = 40012 // 插件注册失败
	ErrPluginUnregisterFailed   = 40013 // 插件注销失败
	ErrPluginInvalidName        = 40014 // 插件名称无效
	ErrPluginInvalidVersion     = 40015 // 插件版本无效
	ErrPluginDependencyMissing  = 40016 // 插件依赖缺失
	ErrPluginPermissionDenied   = 40017 // 插件权限不足
	ErrPluginLoadFailed         = 40018 // 插件加载失败
	ErrPluginUnloadFailed       = 40019 // 插件卸载失败
	ErrPluginStatusInvalid      = 40020 // 插件状态无效
	ErrPluginParamInvalid       = 40021 // 插件参数无效
	ErrPluginCompileFailed      = 40022 // 插件编译失败
	ErrPluginSystemError        = 40023 // 插件系统错误
)

func init() {
	code.Register(ErrPluginNotFound, "plugin not found: {id}")
	code.Register(ErrPluginExists, "plugin already exists: {name}")
	code.Register(ErrPluginNotRunning, "plugin is not running: {id}")
	code.Register(ErrPluginStartFailed, "failed to start plugin: {id}")
	code.Register(ErrPluginStopFailed, "failed to stop plugin: {id}")
	code.Register(ErrPluginInvalidConfig, "invalid plugin config")
	code.Register(ErrPluginTimeout, "plugin operation timeout")
	code.Register(ErrPluginInvalidOpenapi3Doc, "invalid OpenAPI 3.0 document")
	code.Register(ErrPluginGRPCConnFailed, "gRPC connection failed")
	code.Register(ErrPluginInstallFailed, "failed to install plugin: {name}")
	code.Register(ErrPluginUninstallFailed, "failed to uninstall plugin: {name}")
	code.Register(ErrPluginRegisterFailed, "failed to register plugin: {id}")
	code.Register(ErrPluginUnregisterFailed, "failed to unregister plugin: {id}")
	code.Register(ErrPluginInvalidName, "invalid plugin name: {name}")
	code.Register(ErrPluginInvalidVersion, "invalid plugin version: {version}")
	code.Register(ErrPluginDependencyMissing, "plugin dependency missing: {dependency}")
	code.Register(ErrPluginPermissionDenied, "plugin permission denied")
	code.Register(ErrPluginLoadFailed, "failed to load plugin: {path}")
	code.Register(ErrPluginUnloadFailed, "failed to unload plugin: {id}")
	code.Register(ErrPluginStatusInvalid, "invalid plugin status: {status}")
	code.Register(ErrPluginParamInvalid, "invalid plugin parameter")
	code.Register(ErrPluginCompileFailed, "failed to compile plugin: {reason}")
	code.Register(ErrPluginSystemError, "plugin system error: {reason}")
}
