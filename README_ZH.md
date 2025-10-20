go-easyhttp
=================

轻量、易用的 Go HTTP 客户端，内置合理默认值、连接池与超时配置。JSON 采用 `bytedance/sonic` 实现高性能编解码。

[English](README.md) | 中文文档

特性
- 链式请求构建器（GET/POST/PUT/DELETE）
- 每主机连接上限与超时配置
- 高性能 JSON 编解码（`sonic`）
- 简洁 API，极少心智负担

安装
```bash
go get github.com/lockp111/go-easyhttp
```

快速上手
```go
client := easyhttp.NewClient(easyhttp.Config{ /* 合理默认值可直接使用 */ })

req, _ := easyhttp.NewGet("https://httpbin.org/get")
req.AddHeader("Accept", "application/json").AddQuery("q", "ping")

resp, err := client.Fetch(context.Background(), req)
if err != nil { panic(err) }

fmt.Println(resp.Status, resp.GetBodyString())
```

发送 JSON 请求
```go
payload := map[string]any{"name": "alice"}
req, _ := easyhttp.NewPost("https://httpbin.org/post")
_, _ = req.SetJSON(payload) // 自动设置 Content-Type 与 body
```

配置说明
```go
easyhttp.Config{
    MaxConns:        1000,             // 每主机最大连接数
    Timeout:         30 * time.Second, // 总请求超时 >= 响应超时 + 连接超时
    ResponseTimeout: 10 * time.Second, // 响应头超时
    ConnTimeout:     2 * time.Second,  // 建连超时（含 TLS 握手）
    IdleConnTimeout: 90 * time.Second, // 空闲连接回收超时
    DisableHttp2:    false,            // 是否关闭 HTTP/2
}
client := easyhttp.NewClient(cfg)
```

API 概览
- `Client`
  - `Fetch(ctx, *Request) (*Response, error)`
- `Request`
  - 构造：`NewRequest`、`NewGet`、`NewPost`、`NewPut`、`NewDelete`
  - 构建：`SetPath`、`AddHeader`、`SetHeader`、`DelHeader`、`AddQuery`、`SetQuery`、`DelQuery`、`SetBody`、`SetJSON`
  - 访问：`GetUrl`、`GetMethod`、`GetHeader`、`GetBody` 等
- `Response`
  - `GetBodyBytes`、`GetBodyString`、`Unmarshal`

许可证
MIT


