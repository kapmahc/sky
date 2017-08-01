package auth

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/sky/web"
	"github.com/kapmahc/sky/web/cache"
	"github.com/kapmahc/sky/web/i18n"
	"github.com/kapmahc/sky/web/job"
	"github.com/kapmahc/sky/web/security"
	"github.com/kapmahc/sky/web/settings"
	"github.com/kapmahc/sky/web/uploader"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Dao                  *Dao                  `inject:""`
	Jwt                  *Jwt                  `inject:""`
	Db                   *gorm.DB              `inject:""`
	I18n                 *i18n.I18n            `inject:""`
	Cache                *cache.Cache          `inject:""`
	Render               *render.Render        `inject:""`
	Uploader             uploader.Store        `inject:""`
	Hmac                 *security.Hmac        `inject:""`
	Server               *job.Server           `inject:""`
	Settings             *settings.Settings    `inject:""`
	Wrapper              *web.Wrapper          `inject:""`
	MustSignInMiddleware *MustSignInMiddleware `inject:""`
}

// Open open beans
func (p *Plugin) Open(*inject.Graph) error {
	return nil
}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

func init() {
	web.Register(&Plugin{})
}
