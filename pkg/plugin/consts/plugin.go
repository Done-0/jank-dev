// Package consts 提供应用程序常量定义
// 创建者：Done-0
// 创建时间：2025-08-05
package consts

const (
	// 插件目录和文件
	PluginDir        = "plugins"     // 插件目录
	PluginConfigFile = "plugin.json" // 插件配置文件
	PluginBinDir     = "bin"         // 插件二进制文件目录
	PluginMainFile   = "main.go"     // 插件主文件名
)

const (
	// 插件状态
	PluginStatusReady      = "ready"       // 插件就绪
	PluginStatusLoaded     = "loaded"      // 插件已加载
	PluginStatusRunning    = "running"     // 插件正在运行
	PluginStatusStopped    = "stopped"     // 插件已停止
	PluginStatusError      = "error"       // 插件错误
	PluginStatusAvailable  = "available"   // 未注册但可用（有二进制文件）
	PluginStatusSourceOnly = "source_only" // 未注册但有源码（无二进制文件）
	PluginStatusIncomplete = "incomplete"  // 未注册且不完整（既无二进制也无源码）
)

const (
	// 插件类型标识符（按功能分类）
	PluginTypeProvider = "provider" // 数据提供者插件
	PluginTypeFilter   = "filter"   // 数据过滤插件
	PluginTypeHandler  = "handler"  // 业务处理插件
	PluginTypeNotifier = "notifier" // 通知插件
)

const (
	// Go编译相关常量
	GoCommand      = "go"            // Go命令
	GoBuildCommand = "build"         // Go build 子命令
	CGODisabledEnv = "CGO_ENABLED=0" // 禁用 CGO 环境变量
)

const (
	// 插件执行相关常量
	PluginPayloadKey = "payload" // 插件执行时的 payload 键名

	// 插件服务标识符
	PluginServiceKey = "jank-plugin" // 标准插件服务标识符
)

var (
	// Go mod tidy 参数
	GoModTidyArgs = []string{"mod", "tidy"} // go mod tidy 参数
)
