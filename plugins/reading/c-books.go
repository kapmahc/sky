package reading

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/epub"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) indexBooks(c *axe.Context) (interface{}, error) {

	var total int64
	if err := p.Db.Model(&Book{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pag := web.NewPagination(c.Request, total)

	var books []Book
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&books).Error; err != nil {
		return nil, err
	}
	for _, b := range books {
		pag.Items = append(pag.Items, b)
	}
	c.JSON(http.StatusOK, pag)
	return nil
}

func (p *Plugin) showBook(c *axe.Context) (interface{}, error) {
	data := axe.H{}

	var buf bytes.Buffer
	it, bk, err := p.readBook(c.Params["id"])
	if err != nil {
		return nil, err
	}
	var notes []Note
	if err := p.Db.Order("updated_at DESC").Find(&notes).Error; err != nil {
		return nil, err
	}
	data["notes"] = notes
	// c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.writePoints(
		&buf,
		fmt.Sprintf("%s/reading/pages/%d", web.Backend(), it.ID),
		bk.Ncx.Points,
	)
	data["homeage"] = buf.String()
	data["book"] = it

	c.JSON(http.StatusOK, data)
	return nil
}

func (p *Plugin) showPage(c *axe.Context) {
	if err := p.readBookPage(c.Writer, c.Params["id"], c.Params["href"][1:]); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// -----------------------

func (p *Plugin) readBookPage(w http.ResponseWriter, id string, name string) error {
	_, bk, err := p.readBook(id)
	if err != nil {
		return nil, err
	}
	for _, fn := range bk.Files() {
		if strings.HasSuffix(fn, name) {
			for _, mf := range bk.Opf.Manifest {
				if mf.Href == name {
					rdr, err := bk.Open(name)
					if err != nil {
						return nil, err
					}
					defer rdr.Close()
					body, err := ioutil.ReadAll(rdr)
					if err != nil {
						return nil, err
					}
					w.Header().Set("Content-Type", mf.MediaType)
					w.Write(body)
					return nil
				}
			}
		}
	}
	return nil, errors.New("not found")
}

func (p *Plugin) writePoints(wrt io.Writer, href string, points []epub.NavPoint) {
	wrt.Write([]byte("<ol>"))
	for _, it := range points {
		wrt.Write([]byte("<li>"))
		fmt.Fprintf(
			wrt,
			`<a href="%s/%s" target="_blank">%s</a>`,
			href,
			it.Content.Src,
			it.Text,
		)
		p.writePoints(wrt, href, it.Points)
		wrt.Write([]byte("</li>"))
	}
	wrt.Write([]byte("</ol>"))
}

func (p *Plugin) readBook(id string) (*Book, *epub.Book, error) {
	var book Book
	if err := p.Db.
		Where("id = ?", id).First(&book).Error; err != nil {
		return nil, nil, err
	}
	bk, err := epub.Open(path.Join(p.root(), book.File))
	return &book, bk, err
}

func (p *Plugin) destroyBook(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Book{}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil

}
