package easyhttp

import (
	"net/http"

	"github.com/bytedance/sonic"
)

type Response struct {
	Status     int
	StatusText string
	Header     http.Header
	body       []byte
}

func (r *Response) GetBodyBytes() []byte {
	return r.body
}

func (r *Response) GetBodyString() string {
	return string(r.body)
}

func (r *Response) Unmarshal(v any) error {
	return sonic.Unmarshal(r.body, v)
}
