package i18n

import (
	"context"
	"net/http"

	"github.com/kapmahc/sky/web"
	"github.com/urfave/negroni"
	"golang.org/x/text/language"
)

const (
	// LOCALE locale key
	LOCALE = web.K("locale")
)

// Middleware detect language from http request
func (p *I18n) Middleware() (negroni.HandlerFunc, error) {
	langs, err := p.Store.Languages()
	if err != nil {
		return nil, err
	}
	var tags []language.Tag
	for _, l := range langs {
		tags = append(tags, language.Make(l))
	}
	matcher := language.NewMatcher(tags)

	return func(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		const key = string(LOCALE)
		lang, write := p.detect(key, req)
		tag, _, _ := matcher.Match(language.Make(lang))
		lang = tag.String()
		if write {
			http.SetCookie(wrt, &http.Cookie{
				Name:   key,
				Value:  lang,
				MaxAge: 1<<32 - 1,
			})
		}
		ctx := context.WithValue(req.Context(), LOCALE, lang)
		ctx = context.WithValue(ctx, web.K("languages"), langs)
		next(wrt, req.WithContext(ctx))
	}, nil
}

func (p *I18n) detect(k string, r *http.Request) (string, bool) {
	// 1. Check URL arguments.
	if lang := r.URL.Query().Get(k); lang != "" {
		return lang, true
	}

	// 2. Get language information from cookies.
	if ck, er := r.Cookie(k); er == nil {
		return ck.Value, false
	}

	// 3. Get language information from 'Accept-Language'.
	if al := r.Header.Get("Accept-Language"); len(al) > 4 {
		return al[:5], true // Only compare first 5 letters.
	}

	return "", true
}
