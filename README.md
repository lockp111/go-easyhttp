go-easyhttp
=================

Lightweight, ergonomic HTTP client for Go, with sane defaults, connection pooling, per-host limits, and timeouts. JSON is powered by `bytedance/sonic` for high performance.

English | [中文文档](README_ZH.md)

Features
- Chainable request builder (GET/POST/PUT/DELETE)
- Per-host connection limits and timeouts
- Fast JSON marshal/unmarshal with `sonic`
- Simple, minimal API surface
 - Pluggable options (custom http.Client/Transport, timeouts, conn limits)

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
    client := easyhttp.NewClient(
        easyhttp.Config{
            MaxConns:        100,
            Timeout:         5 * time.Second,
            ResponseTimeout: 4 * time.Second,
            ConnTimeout:     1 * time.Second,
            IdleConnTimeout: 90 * time.Second,
        },
        // Options are optional and can override parts of the config at runtime
        easyhttp.WithTimeout(10*time.Second),
        easyhttp.WithMaxConns(200),
    )

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

Streaming request body
```go
// Stream from an io.Reader without buffering into memory
file, _ := os.Open("bigfile.bin")
defer file.Close()

req, _ := easyhttp.NewPost("https://httpbin.org/post")
req.SetHeader("Content-Type", "application/octet-stream")
req.SetBodyReader(file)
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
  - Options: `WithHTTPClient`, `WithTransport`, `WithTimeout`, `WithMaxConns`, `WithConnTimeout`, `WithResponseHeaderTimeout`, `WithIdleConnTimeout`, `WithDisableHTTP2`
- type `Request`
  - Constructors: `NewRequest`, `NewGet`, `NewPost`, `NewPut`, `NewDelete`, `NewPatch`, `NewHead`, `NewOptions`, `NewConnect`, `NewTrace`
  - Builders: `SetPath`, `AddHeader`, `SetHeader`, `DelHeader`, `AddQuery`, `SetQuery`, `DelQuery`, `SetBody`, `SetJSON`, `SetBodyReader`
  - Getters: `GetUrl`, `GetMethod`, `GetHeader`, `GetBody`, etc.
- type `Response`
  - `GetBodyBytes`, `GetBodyString`, `Unmarshal`

License
MIT


Further reading
- Best practices: [docs/best-practices.md](docs/best-practices.md)
- Examples: see the `examples/` directory. Each example is an isolated Go module with its own `go.mod` (using `replace`), so the root `go.mod` remains untouched.


