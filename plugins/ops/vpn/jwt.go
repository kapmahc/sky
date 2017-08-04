package vpn

import (
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) generateToken(years int) ([]byte, error) {
	now := time.Now()
	cm := jws.Claims{}
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(years, 0, 0))
	cm.Set("act", "vpn")

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

func (p *Plugin) tokenMiddleware(c *axe.Context) error {
	lng := c.Payload[i18n.LOCALE].(string)
	tk, err := jws.ParseJWTFromRequest(c.Request)
	if err != nil {
		return err
	}
	if err := tk.Validate(p.Key, p.Method); err != nil {
		return err
	}

	if act := tk.Claims().Get("act"); act != nil && act.(string) == "vpn" {
		return nil
	}
	return p.I18n.E(lng, "errors.forbidden")
}
