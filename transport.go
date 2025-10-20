package easyhttp

import (
	"context"
	"net"
	"net/http"
	"time"
)

type Config struct {
	MaxConns        int           // Maximum concurrent connections per host
	Timeout         time.Duration // Total timeout, Timeout >= ResponseTimeout + ConnTimeout
	ResponseTimeout time.Duration // Response header timeout
	ConnTimeout     time.Duration // Connection timeout
	IdleConnTimeout time.Duration // Idle connection timeout
	DisableHttp2    bool          // Disable HTTP/2 when true
}

func newTransport(cfg Config) *http.Transport {
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// Connection timeout setup
			c, err := net.DialTimeout(network, addr, cfg.ConnTimeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TLSHandshakeTimeout:   cfg.ConnTimeout,     // TLS handshake timeout
		ResponseHeaderTimeout: cfg.ResponseTimeout, // Response header timeout
		ForceAttemptHTTP2:     !cfg.DisableHttp2,   // Enable HTTP/2
		MaxIdleConns:          0,                   // Max keep-alive connections, unlimited
		MaxIdleConnsPerHost:   cfg.MaxConns,        // Keep-alive connections per host
		MaxConnsPerHost:       cfg.MaxConns,        // Max total connections per host
		IdleConnTimeout:       cfg.IdleConnTimeout, // Idle connection timeout
	}
}
