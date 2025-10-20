最佳实践
========

本文总结使用 `go-easyhttp` 时的推荐做法。

客户端生命周期
- 创建并复用一个长期存活的 `Client`，以充分利用连接池。
- 通过 `Config` 设置全局超时，再用 Option 进行细化。

超时建议
- 保证 `Timeout >= ConnTimeout + ResponseTimeout`。
- 对低延迟请求，可在构造时用 `WithTimeout` 单独调整。

请求头与 JSON
- 发送 JSON 时使用 `SetJSON`，自动设置 `Content-Type: application/json`。
- 多值使用 `AddHeader`，单值使用 `SetHeader`。

流式请求体
- 使用 `SetBodyReader(io.Reader)` 上传大文件，避免一次性占用内存。
- 记得为流式上传设置合适的 `Content-Type`。

上下文使用
- 通过 `Fetch` 传入 `context.Context` 以支持每次请求的取消。
- 可搭配 `context.WithTimeout` 为单次请求设置临时截止时间。

传输层调优
- 使用 `WithMaxConns`、`WithConnTimeout`、`WithIdleConnTimeout` 等 Option 调整连接行为。
- 只有在上游要求时才关闭 HTTP/2（`WithDisableHTTP2(true)`）。

错误处理
- 先检查 `Fetch` 返回的 `error`。若非空则请求未成功完成。
- 再根据 `Response.Status` 判断是否解析/处理响应体。

示例
- 参见 `examples/` 目录中的示例代码。每个示例都是独立 Go 模块，不会修改根目录 `go.mod`。


