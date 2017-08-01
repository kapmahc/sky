package web

import (
	"net/http"
)

// HandlerFunc handler func
type HandlerFunc func(c *Context)

// Router http router
type Router struct {
	root   string
	routes []route
}

// Group add group
func (p *Router) Group(path string, router *Router) {

}

// Get http get
func (p *Router) Get(path string, handlers ...HandlerFunc) {
	p.add(http.MethodGet, path, handlers...)
}

// Post http post
func (p *Router) Post(path string, handlers ...HandlerFunc) {
	p.add(http.MethodPost, path, handlers...)
}

// Patch http patch
func (p *Router) Patch(path string, handlers ...HandlerFunc) {
	p.add(http.MethodPatch, path, handlers...)
}

// Delete http delete
func (p *Router) Delete(path string, handlers ...HandlerFunc) {
	p.add(http.MethodDelete, path, handlers...)
}

func (p *Router) add(method, path string, handlers ...HandlerFunc) {
	p.routes = append(p.routes, route{
		method:   method,
		path:     path,
		handlers: handlers,
	})
}

type route struct {
	method   string
	path     string
	handlers []HandlerFunc
}
