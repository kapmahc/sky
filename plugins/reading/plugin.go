package reading

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/axe/job"
	"github.com/kapmahc/sky/plugins/auth"
	"github.com/kapmahc/sky/web"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db   *gorm.DB   `inject:""`
	I18n *i18n.I18n `inject:""`
	Jwt  *auth.Jwt  `inject:""`
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

// Workers background workers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

func init() {
	web.Register(&Plugin{})
}
