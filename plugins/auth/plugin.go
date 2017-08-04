package auth

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/axe/cache"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/axe/job"
	"github.com/kapmahc/axe/security"
	"github.com/kapmahc/axe/settings"
	"github.com/kapmahc/axe/uploader"
	"github.com/kapmahc/sky/web"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Dao      *Dao               `inject:""`
	Jwt      *Jwt               `inject:""`
	Db       *gorm.DB           `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Cache    *cache.Cache       `inject:""`
	Uploader uploader.Store     `inject:""`
	Hmac     *security.Hmac     `inject:""`
	Server   *job.Server        `inject:""`
	Settings *settings.Settings `inject:""`
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
