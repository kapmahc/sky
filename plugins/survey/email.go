package survey

import (
	"fmt"

	"github.com/kapmahc/axe/job"
	"github.com/kapmahc/sky/plugins/auth"
	"github.com/kapmahc/sky/web"
	log "github.com/sirupsen/logrus"
)

const (
	actApply  = "apply"
	actCancel = "cancel"
)

func (p *Plugin) _sendEmail(lng string, form *Form, record *Record, act string) {

	obj := struct {
		Home   string
		Title  string
		Apply  string
		Cancel string
	}{
		Home:   web.Home(),
		Title:  form.Title,
		Apply:  fmt.Sprintf("/forms/apply/%d", form.ID),
		Cancel: fmt.Sprintf("/forms/cancel/%d", form.ID),
	}
	subject, err := p.I18n.F(lng, fmt.Sprintf("forms.emails.%s.subject", act), obj)
	if err != nil {
		log.Error(err)
		return
	}
	body, err := p.I18n.F(lng, fmt.Sprintf("forms.emails.%s.body", act), obj)
	if err != nil {
		log.Error(err)
		return
	}

	// -----------------------
	p.Server.Send(job.PriorityLow, auth.SendEmailJob, map[string]string{
		"to":      record.Email,
		"subject": subject,
		"body":    body,
	})
}
