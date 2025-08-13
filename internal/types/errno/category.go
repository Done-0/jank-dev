// Package errno 分类模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-13
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 分类模块错误码: 60000 ~ 69999
const (
	ErrCategoryCreateFailed = 60001 // 创建分类失败
	ErrCategoryGetFailed    = 60002 // 获取分类失败
	ErrCategoryUpdateFailed = 60003 // 更新分类失败
	ErrCategoryDeleteFailed = 60004 // 删除分类失败
	ErrCategoryListFailed   = 60005 // 获取分类列表失败
)

func init() {
	code.Register(ErrCategoryCreateFailed, "create category failed: {name}")
	code.Register(ErrCategoryGetFailed, "get category failed: {id}")
	code.Register(ErrCategoryUpdateFailed, "update category failed: {id}")
	code.Register(ErrCategoryDeleteFailed, "delete category failed: {id}")
	code.Register(ErrCategoryListFailed, "list categories failed: {msg}")
}
