package survey

import (
	"errors"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/axe"
)

const (
	actApply  = "apply"
	actCancel = "cancel"
)

func (p *Plugin) generateToken(fid, act string, begin, end time.Time) ([]byte, error) {
	cm := jws.Claims{}
	cm.SetNotBefore(begin)
	cm.SetExpiration(end)
	cm.Set("act", act)
	cm.Set("form", fid)

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

func (p *Plugin) parseToken(c *axe.Context, act string) (*Form, error) {
	tk, err := jws.ParseJWTFromRequest(c.Request)
	if err != nil {
		return nil, err
	}
	if err := tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}

	if tk.Claims().Get("act").(string) != act {
		return nil, errors.New("bad action")
	}

	var fm Form
	if err := p.Db.Where("uid = ?", tk.Claims().Get("form")).First(&fm).Error; err != nil {
		return nil, err
	}
	return &fm, nil
}
