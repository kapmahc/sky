package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/web"
)

const (
	// TOKEN token session key
	TOKEN = "token"
	// UID uid key
	UID = "uid"
	// CurrentUser current-user key
	CurrentUser = web.K("currentUser")
	// IsAdmin is-admin key
	IsAdmin = web.K("isAdmin")
)

//Jwt jwt helper
type Jwt struct {
	Key     []byte               `inject:"jwt.key"`
	Method  crypto.SigningMethod `inject:"jwt.method"`
	Dao     *Dao                 `inject:""`
	I18n    *i18n.I18n           `inject:""`
	Wrapper *web.Wrapper         `inject:""`
}

//Validate check jwt
func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

func (p *Jwt) parse(r *http.Request) (jwt.Claims, error) {
	tk, err := jws.ParseJWTFromRequest(r)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

//Sum create jwt token
func (p *Jwt) Sum(cm jws.Claims, exp time.Duration) ([]byte, error) {
	kid := uuid.New().String()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.Add(exp))
	cm.Set("kid", kid)
	//TODO using kid

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

func (p *Jwt) getUserFromRequest(c *web.Context) (*User, error) {
	lng := c.Get(i18n.LOCALE).(string)
	cm, err := p.parse(c.Request)
	if err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByUID(cm.Get(UID).(string))
	if err != nil {
		return nil, err
	}
	if !user.IsConfirm() {
		return nil, p.I18n.E(lng, "auth.errors.user.not-confirm")
	}
	if user.IsLock() {
		return nil, p.I18n.E(lng, "auth.errors.user.is-lock")
	}
	return user, nil
}

// CurrentUserMiddleware current-user middleware
func (p *Jwt) CurrentUserMiddleware(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if user, err := p.getUserFromRequest(p.Wrapper.Context(wrt, req)); err == nil {
		ctx := req.Context()
		ctx = context.WithValue(ctx, CurrentUser, user)
		ctx = context.WithValue(ctx, IsAdmin, p.Dao.Is(user.ID, RoleAdmin))
		next(wrt, req.WithContext(ctx))
		return
	}
	next(wrt, req)
}
