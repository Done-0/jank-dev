// Package errno 插件模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-06
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 插件模块错误码: 30000 ~ 39999
const (
	ErrPluginNotFound         = 30001 // 插件不存在
	ErrPluginParamInvalid     = 30002 // 插件参数无效
	ErrPluginRegisterFailed   = 30003 // 插件注册失败
	ErrPluginUnregisterFailed = 30004 // 插件注销失败
)

func init() {
	code.Register(ErrPluginNotFound, "plugin not found: {plugin_id}")
	code.Register(ErrPluginParamInvalid, "invalid plugin parameter: {msg}")
	code.Register(ErrPluginRegisterFailed, "failed to register plugin: {plugin_id}")
	code.Register(ErrPluginUnregisterFailed, "failed to unregister plugin: {plugin_id}")
}
