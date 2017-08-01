package web

import (
	"net"
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/go-playground/form"
)

// H hash
type H map[string]interface{}

// K key
type K string

// Context http context
type Context struct {
	Params  map[string]string
	Request *http.Request
	Writer  http.ResponseWriter

	decoder  *form.Decoder
	validate *validator.Validate
}

// Get get value from request context
func (p *Context) Get(k K) interface{} {
	return p.Request.Context().Value(k)
}

// Bind bind json form
func (p *Context) Bind(f interface{}) error {
	if e := p.Request.ParseForm(); e != nil {
		return e
	}
	if e := p.decoder.Decode(f, p.Request.Form); e != nil {
		return e
	}
	return p.validate.Struct(f)
}

// ClientIP http client ip
func (p *Context) ClientIP() string {
	ip := p.Request.Header.Get("X-Forwarded-For")
	if idx := strings.IndexByte(ip, ','); idx >= 0 {
		ip = ip[0:idx]
	}
	ip = strings.TrimSpace(ip)
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(p.Request.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip = p.Request.Header.Get("X-Appengine-Remote-Addr"); ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(p.Request.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
