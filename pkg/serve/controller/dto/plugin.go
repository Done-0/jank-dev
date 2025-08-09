// Package dto 提供插件相关的数据传输对象定义
package dto

// RegisterPluginRequest 注册插件请求
type RegisterPluginRequest struct {
	ID string `json:"id" validate:"required,min=1,max=100"` // 插件ID
}

// UnregisterPluginRequest 注销插件请求
type UnregisterPluginRequest struct {
	ID string `json:"id" validate:"required"` // 插件ID
}

// GetPluginRequest 获取插件信息请求
type GetPluginRequest struct {
	ID string `query:"id" validate:"required"` // 插件ID
}

// ListPluginsRequest 列举插件请求
type ListPluginsRequest struct {
	Status   string `query:"status" validate:"omitempty"`                  // 插件状态筛选
	PageNo   int64  `query:"page_no" validate:"omitempty,min=1"`           // 页码
	PageSize int64  `query:"page_size" validate:"omitempty,min=1,max=100"` // 每页数量
}

// ExecutePluginRequest 执行插件方法请求
type ExecutePluginRequest struct {
	ID     string         `json:"id" validate:"required"`     // 插件ID
	Method string         `json:"method" validate:"required"` // 方法名
	Args   map[string]any `json:"args" validate:"omitempty"`  // 方法参数
}

// StartPluginRequest 启动插件请求
type StartPluginRequest struct {
	ID string `json:"id" validate:"required"` // 插件ID
}

// StopPluginRequest 停止插件请求
type StopPluginRequest struct {
	ID string `json:"id" validate:"required"` // 插件ID
}

// HeartbeatRequest 心跳请求 (兼容旧controller)
type HeartbeatRequest struct {
	PluginId string            `json:"plugin_id" validate:"required"` // 插件ID
	Status   int64             `json:"status" validate:"omitempty"`   // 当前状态
	Metadata map[string]string `json:"metadata" validate:"omitempty"` // 元数据信息
}

// UpdatePluginStatusRequest 更新插件状态请求 (兼容旧controller)
type UpdatePluginStatusRequest struct {
	PluginId string `json:"plugin_id" validate:"required"`    // 插件ID
	Status   int64  `json:"status" validate:"required,min=0"` // 新状态
}
