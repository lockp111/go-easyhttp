package easyhttp

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/bytedance/sonic"
)

const (
	MethodGet     = http.MethodGet
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodDelete  = http.MethodDelete
	MethodPatch   = http.MethodPatch
	MethodHead    = http.MethodHead
	MethodOptions = http.MethodOptions
	MethodConnect = http.MethodConnect
	MethodTrace   = http.MethodTrace
)

type Request struct {
	method string
	urls   *url.URL
	header http.Header
	query  url.Values
	body   *bytes.Buffer
}

func NewRequest(method, baseUrl string) (*Request, error) {
	urls, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}

	return &Request{
		method: method,
		urls:   urls,
		header: make(http.Header),
		query:  make(url.Values),
	}, nil
}

func NewGet(baseUrl string) (*Request, error) {
	return NewRequest(MethodGet, baseUrl)
}

func NewPost(baseUrl string) (*Request, error) {
	return NewRequest(MethodPost, baseUrl)
}

func NewPut(baseUrl string) (*Request, error) {
	return NewRequest(MethodPut, baseUrl)
}

func NewDelete(baseUrl string) (*Request, error) {
	return NewRequest(MethodDelete, baseUrl)
}

func (r *Request) AddHeader(key, value string) *Request {
	r.header.Add(key, value)
	return r
}

func (r *Request) SetHeader(key, value string) *Request {
	r.header.Set(key, value)
	return r
}

func (r *Request) DelHeader(key string) *Request {
	r.header.Del(key)
	return r
}

func (r *Request) DelQuery(key string) *Request {
	r.query.Del(key)
	return r
}

func (r *Request) AddQuery(key, value string) *Request {
	r.query.Add(key, value)
	return r
}

func (r *Request) SetQuery(key, value string) *Request {
	r.query.Set(key, value)
	return r
}

func (r *Request) SetPath(path string) *Request {
	r.urls.Path = path
	return r
}

func (r *Request) SetBody(body []byte) *Request {
	r.body = bytes.NewBuffer(body)
	return r
}

func (r *Request) SetJSON(body any) (*Request, error) {
	data, err := sonic.Marshal(body)
	if err != nil {
		return nil, err
	}
	r.body = bytes.NewBuffer(data)
	r.header.Set("Content-Type", "application/json")
	return r, nil
}

func (r *Request) GetBaseUrl() string {
	return r.urls.Scheme + "://" + r.urls.Host
}

func (r *Request) GetMethod() string {
	return r.method
}

func (r *Request) GetPath() string {
	return r.urls.EscapedPath()
}

func (r *Request) GetUrl() string {
	return r.urls.String()
}

func (r *Request) GetRawQuery() url.Values {
	return r.query
}

func (r *Request) GetQuery() string {
	return r.query.Encode()
}

func (r *Request) GetHeader() http.Header {
	return r.header
}

func (r *Request) GetBody() *bytes.Buffer {
	return r.body
}
