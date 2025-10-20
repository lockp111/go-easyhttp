package easyhttp

import (
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
	d := &net.Dialer{
		Timeout:   cfg.ConnTimeout,
		KeepAlive: 30 * time.Second,
	}
	return &http.Transport{
		DialContext:           d.DialContext,
		TLSHandshakeTimeout:   cfg.ConnTimeout,     // TLS handshake timeout
		ResponseHeaderTimeout: cfg.ResponseTimeout, // Response header timeout
		ExpectContinueTimeout: 1 * time.Second,     // Reasonable default
		ForceAttemptHTTP2:     !cfg.DisableHttp2,   // Enable HTTP/2 unless disabled
		MaxIdleConns:          0,                   // Max keep-alive connections, unlimited
		MaxIdleConnsPerHost:   cfg.MaxConns,        // Keep-alive connections per host
		MaxConnsPerHost:       cfg.MaxConns,        // Max total connections per host
		IdleConnTimeout:       cfg.IdleConnTimeout, // Idle connection timeout
	}
}
