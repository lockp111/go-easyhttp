package easyhttp

import (
	"context"
	"io"
	"net/http"
	"time"
)

type Option func(*Client)

type Client struct {
	Conf Config
	conn *http.Client
}

func NewClient(cfg Config, opts ...Option) *Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = time.Second * 30
	}
	if cfg.MaxConns == 0 {
		cfg.MaxConns = 1000
	}
	if cfg.IdleConnTimeout == 0 {
		cfg.IdleConnTimeout = time.Second * 90
	}
	// Provide reasonable defaults for connection and response header timeouts
	if cfg.ConnTimeout == 0 {
		cfg.ConnTimeout = 5 * time.Second
	}
	if cfg.ResponseTimeout == 0 && cfg.Timeout > cfg.ConnTimeout {
		cfg.ResponseTimeout = cfg.Timeout - cfg.ConnTimeout
	}

	client := &Client{
		Conf: cfg,
		conn: &http.Client{
			// cfg.Timeout is already a time.Duration
			Timeout:   cfg.Timeout,
			Transport: newTransport(cfg),
		},
	}
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// WithHTTPClient overrides the underlying http.Client
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.conn = hc
		}
	}
}

// WithTransport overrides the transport on the underlying http.Client
func WithTransport(tr http.RoundTripper) Option {
	return func(c *Client) {
		if tr != nil {
			c.conn.Transport = tr
		}
	}
}

// WithTimeout sets the overall client timeout
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		if d > 0 {
			c.Conf.Timeout = d
			c.conn.Transport = newTransport(c.Conf)
		}
	}
}

// WithMaxConns sets per-host connection limits
func WithMaxConns(n int) Option {
	return func(c *Client) {
		if n > 0 {
			c.Conf.MaxConns = n
			// Rebuild transport to apply
			c.conn.Transport = newTransport(c.Conf)
		}
	}
}

// WithConnTimeout sets TCP connect timeout and TLS handshake timeout
func WithConnTimeout(d time.Duration) Option {
	return func(c *Client) {
		if d > 0 {
			c.Conf.ConnTimeout = d
			c.conn.Transport = newTransport(c.Conf)
		}
	}
}

// WithResponseHeaderTimeout sets response header timeout
func WithResponseHeaderTimeout(d time.Duration) Option {
	return func(c *Client) {
		if d > 0 {
			c.Conf.ResponseTimeout = d
			c.conn.Transport = newTransport(c.Conf)
		}
	}
}

// WithIdleConnTimeout sets idle connection timeout
func WithIdleConnTimeout(d time.Duration) Option {
	return func(c *Client) {
		if d > 0 {
			c.Conf.IdleConnTimeout = d
			c.conn.Transport = newTransport(c.Conf)
		}
	}
}

// WithDisableHTTP2 disables HTTP/2 when set to true
func WithDisableHTTP2(disable bool) Option {
	return func(c *Client) {
		c.Conf.DisableHttp2 = disable
		c.conn.Transport = newTransport(c.Conf)
	}
}

func (r *Client) Fetch(ctx context.Context, req *Request) (*Response, error) {
	req.urls.RawQuery = req.GetQuery()
	// If not explicitly using nil, prepare body like this to avoid panic
	var reqBody io.Reader = http.NoBody
	if req.body != nil {
		reqBody = req.GetBody()
	}
	httpReq, err := http.NewRequestWithContext(ctx, req.GetMethod(), req.GetUrl(), reqBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header = req.GetHeader()
	resp, err := r.conn.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Status:     resp.StatusCode,
		StatusText: resp.Status,
		Header:     resp.Header,
		body:       respBody,
	}, nil
}
