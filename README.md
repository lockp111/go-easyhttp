go-easyhttp
=================

Lightweight, ergonomic HTTP client for Go, with sane defaults, connection pooling, per-host limits, and timeouts. JSON is powered by `bytedance/sonic` for high performance.

English | [中文文档](README_ZH.md)

Features
- Chainable request builder (GET/POST/PUT/DELETE)
- Per-host connection limits and timeouts
- Fast JSON marshal/unmarshal with `sonic`
- Simple, minimal API surface

Install
```bash
go get github.com/lockp111/go-easyhttp
```

Quick start
```go
package main

import (
    "context"
    "fmt"
    "time"

    easyhttp "github.com/lockp111/go-easyhttp"
)

type Resp struct {
    Message string `json:"message"`
}

func main() {
    client := easyhttp.NewClient(easyhttp.Config{
        MaxConns:        100,
        Timeout:         5 * time.Second,
        ResponseTimeout: 4 * time.Second,
        ConnTimeout:     1 * time.Second,
        IdleConnTimeout: 90 * time.Second,
    })

    req, _ := easyhttp.NewGet("https://httpbin.org/get")
    req.AddHeader("Accept", "application/json").AddQuery("q", "ping")

    resp, err := client.Fetch(context.Background(), req)
    if err != nil { panic(err) }

    var data map[string]any
    if err := resp.Unmarshal(&data); err != nil { panic(err) }
    fmt.Println("status:", resp.Status, "args:", data["args"])
}
```

JSON request body
```go
payload := map[string]any{"name": "alice"}
req, _ := easyhttp.NewPost("https://httpbin.org/post")
_, _ = req.SetJSON(payload) // sets Content-Type and body
```

Configuration
```go
cfg := easyhttp.Config{
    MaxConns:        1000,            // per-host limits
    Timeout:         30 * time.Second, // >= ResponseTimeout + ConnTimeout
    ResponseTimeout: 10 * time.Second,
    ConnTimeout:     2 * time.Second,
    IdleConnTimeout: 90 * time.Second,
    DisableHttp2:    false,
}
client := easyhttp.NewClient(cfg)
```

API overview
- type `Client`
  - `Fetch(ctx, *Request) (*Response, error)`
- type `Request`
  - Constructors: `NewRequest`, `NewGet`, `NewPost`, `NewPut`, `NewDelete`
  - Builders: `SetPath`, `AddHeader`, `SetHeader`, `DelHeader`, `AddQuery`, `SetQuery`, `DelQuery`, `SetBody`, `SetJSON`
  - Getters: `GetUrl`, `GetMethod`, `GetHeader`, `GetBody`, etc.
- type `Response`
  - `GetBodyBytes`, `GetBodyString`, `Unmarshal`

License
MIT


