package auth

import (
	"net/http"

	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/web"
)

// MustSignInMiddleware must sign-in
type MustSignInMiddleware struct {
	I18n    *i18n.I18n   `inject:""`
	Wrapper *web.Wrapper `inject:""`
}

func (p *MustSignInMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	c := p.Wrapper.Context(wrt, req)
	lng := c.Get(i18n.LOCALE).(string)

	if user := c.Get(CurrentUser); user == nil {
		http.Error(wrt, p.I18n.T(lng, "auth.errors.please-sign-in"), http.StatusForbidden)
		return
	}
	next(wrt, req)
}

// MustAdminMiddleware must has admin role
type MustAdminMiddleware struct {
	I18n    *i18n.I18n   `inject:""`
	Wrapper *web.Wrapper `inject:""`
}

func (p *MustAdminMiddleware) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	c := p.Wrapper.Context(wrt, req)
	lng := c.Get(i18n.LOCALE).(string)

	if ok := c.Get(IsAdmin); ok == nil || !ok.(bool) {
		http.Error(wrt, p.I18n.T(lng, "errors.not-allow"), http.StatusForbidden)
		return
	}
	next(wrt, req)
}
