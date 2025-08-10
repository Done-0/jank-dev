// Package consts 提供应用程序常量定义
// 创建者：Done-0
// 创建时间：2025-08-05
package consts

const (
	// 通用状态码
	Success = 0 // 成功状态码

	// HTTP 1xx 信息响应
	StatusContinue           = 100 // 继续
	StatusSwitchingProtocols = 101 // 切换协议
	StatusProcessing         = 102 // 处理中
	StatusEarlyHints         = 103 // 早期提示

	// HTTP 2xx 成功响应
	StatusOK                   = 200 // 成功
	StatusCreated              = 201 // 已创建
	StatusAccepted             = 202 // 已接受
	StatusNonAuthoritativeInfo = 203 // 非权威信息
	StatusNoContent            = 204 // 无内容
	StatusResetContent         = 205 // 重置内容
	StatusPartialContent       = 206 // 部分内容

	// HTTP 3xx 重定向响应
	StatusMultipleChoices   = 300 // 多种选择
	StatusMovedPermanently  = 301 // 永久移动
	StatusFound             = 302 // 找到
	StatusSeeOther          = 303 // 查看其他
	StatusNotModified       = 304 // 未修改
	StatusUseProxy          = 305 // 使用代理
	StatusTemporaryRedirect = 307 // 临时重定向
	StatusPermanentRedirect = 308 // 永久重定向

	// HTTP 4xx 客户端错误响应
	StatusBadRequest                   = 400 // 错误请求
	StatusUnauthorized                 = 401 // 未授权
	StatusPaymentRequired              = 402 // 需要付款
	StatusForbidden                    = 403 // 禁止访问
	StatusNotFound                     = 404 // 未找到
	StatusMethodNotAllowed             = 405 // 方法不允许
	StatusNotAcceptable                = 406 // 不可接受
	StatusProxyAuthRequired            = 407 // 需要代理授权
	StatusRequestTimeout               = 408 // 请求超时
	StatusConflict                     = 409 // 冲突
	StatusGone                         = 410 // 已删除
	StatusLengthRequired               = 411 // 需要长度
	StatusPreconditionFailed           = 412 // 前置条件失败
	StatusRequestEntityTooLarge        = 413 // 请求实体过大
	StatusRequestURITooLong            = 414 // 请求URI过长
	StatusUnsupportedMediaType         = 415 // 不支持的媒体类型
	StatusRequestedRangeNotSatisfiable = 416 // 请求范围不满足
	StatusExpectationFailed            = 417 // 预期失败
	StatusTeapot                       = 418 // 我是茶壶 (RFC 2324)
	StatusMisdirectedRequest           = 421 // 误导请求
	StatusUnprocessableEntity          = 422 // 无法处理的实体
	StatusLocked                       = 423 // 已锁定
	StatusFailedDependency             = 424 // 依赖失败
	StatusTooEarly                     = 425 // 过早
	StatusUpgradeRequired              = 426 // 需要升级
	StatusPreconditionRequired         = 428 // 需要前置条件
	StatusTooManyRequests              = 429 // 请求过多
	StatusRequestHeaderFieldsTooLarge  = 431 // 请求头字段过大
	StatusUnavailableForLegalReasons   = 451 // 因法律原因不可用

	// HTTP 5xx 服务器错误响应
	StatusInternalServerError           = 500 // 内部服务器错误
	StatusNotImplemented                = 501 // 未实现
	StatusBadGateway                    = 502 // 网关错误
	StatusServiceUnavailable            = 503 // 服务不可用
	StatusGatewayTimeout                = 504 // 网关超时
	StatusHTTPVersionNotSupported       = 505 // HTTP 版本不支持
	StatusVariantAlsoNegotiates         = 506 // 变体协商
	StatusInsufficientStorage           = 507 // 存储不足
	StatusLoopDetected                  = 508 // 循环检测
	StatusNotExtended                   = 510 // 未扩展
	StatusNetworkAuthenticationRequired = 511 // 需要网络授权
)
