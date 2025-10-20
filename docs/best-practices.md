Best Practices
==============

This guide summarizes recommended patterns when using `go-easyhttp`.

Client lifecycle
- Create a single long-lived `Client` and reuse it. Connection pooling relies on reuse.
- Prefer setting timeouts globally via `Config` and fine-tune with Options.

Timeouts
- Ensure `Timeout >= ConnTimeout + ResponseTimeout`.
- For latency-sensitive requests, adjust with `WithTimeout` at client construction.

Headers and JSON
- Use `SetJSON` for JSON payloads; it sets `Content-Type: application/json`.
- Prefer `AddHeader` for multi-value headers and `SetHeader` for single-value.

Streaming request bodies
- Use `SetBodyReader(io.Reader)` to avoid buffering large payloads in memory.
- Always set an appropriate `Content-Type` for streaming uploads.

Context usage
- Pass `context.Context` to `Fetch` for per-request cancellation.
- Use `context.WithTimeout` for ad-hoc deadlines in addition to client-level timeout.

Transport tuning
- Use Options like `WithMaxConns`, `WithConnTimeout`, `WithIdleConnTimeout` to tune connection behavior.
- Disable HTTP/2 only if required by the upstream (`WithDisableHTTP2(true)`).

Error handling
- Check `Fetch` errors first; if non-nil, the request did not complete successfully.
- Inspect `Response.Status` for non-2xx and decide whether to parse body.

Examples
- See `examples/` for concrete code samples. Each example is an isolated Go module to avoid modifying the root `go.mod`.


