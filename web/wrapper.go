package web

import (
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/go-playground/form"
	"github.com/gorilla/mux"
)

// NewWrapper new wrapper
func NewWrapper() *Wrapper {
	return &Wrapper{
		decoder:  form.NewDecoder(),
		validate: validator.New(),
	}
}

// Wrapper wrapper
type Wrapper struct {
	decoder  *form.Decoder
	validate *validator.Validate
}

// JSON  json render
func (p *Wrapper) JSON(f func(*Context) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := p.Context(w, r)
		v, e := f(c)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(v)
	}
}

// Context new context
func (p *Wrapper) Context(w http.ResponseWriter, r *http.Request) *Context {

	return &Context{
		Writer:   w,
		Request:  r,
		Params:   mux.Vars(r),
		decoder:  p.decoder,
		validate: p.validate,
	}
}
