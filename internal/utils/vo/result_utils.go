// Package vo 提供通用值对象
// 创建者：Done-0
// 创建时间：2025-08-05
package vo

import (
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
)

// Result 通用 API 响应结果结构体
type Result struct {
	errorx.StatusError `json:",omitempty"` // 错误信息
	Data               any                 `json:"data"`      // 响应数据
	RequestId          any                 `json:"requestId"` // 请求ID
	TimeStamp          any                 `json:"timeStamp"` // 响应时间戳
}

// Success 成功返回
// 参数：
//
//	c: Hertz 上下文
//	data: 响应数据
//
// 返回值：
//
//	Result: 成功响应结果
func Success(c *app.RequestContext, data any) Result {
	return Result{
		Data:      data,
		RequestId: string(c.GetHeader("X-Request-ID")),
		TimeStamp: time.Now().Unix(),
	}
}

// Fail 失败返回
// 参数：
//
//	c: Hertz 上下文
//	data: 错误相关数据
//	err: 错误对象
//
// 返回值：
//
//	Result: 失败响应结果
func Fail(c *app.RequestContext, data any, err error) Result {
	var newBizErr errorx.StatusError
	if ok := errors.As(err, &newBizErr); ok {
		return Result{
			StatusError: newBizErr,
			Data:        data,
			RequestId:   string(c.GetHeader("X-Request-ID")),
			TimeStamp:   time.Now().Unix(),
		}
	}

	return Result{
		StatusError: errorx.New(errno.ErrInternalServer).(errorx.StatusError),
		Data:        data,
		RequestId:   string(c.GetHeader("X-Request-ID")),
		TimeStamp:   time.Now().Unix(),
	}
}
