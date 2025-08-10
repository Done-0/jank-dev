// Package dto 提供主题相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-09
package dto

// SwitchThemeRequest 切换主题请求
type SwitchThemeRequest struct {
	ID      string `json:"id" validate:"required"` // 主题ID
	Rebuild bool   `json:"rebuild,omitempty"`      // 强制重新构建
}

// ListThemesRequest 列举主题请求
type ListThemesRequest struct {
	Status   string `query:"status" validate:"omitempty"`                  // 主题状态筛选
	PageNo   int64  `query:"page_no" validate:"omitempty,min=1"`           // 页码
	PageSize int64  `query:"page_size" validate:"omitempty,min=1,max=100"` // 每页数量
}
