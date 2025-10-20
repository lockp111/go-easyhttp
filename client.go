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

	client := &Client{
		conn: &http.Client{
			Timeout:   time.Second * time.Duration(cfg.Timeout),
			Transport: newTransport(cfg),
		},
	}
	for _, opt := range opts {
		opt(client)
	}

	return client
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
