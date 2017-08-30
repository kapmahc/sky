package survey

import (
	"github.com/kapmahc/sky/web"
)

// Form form
type Form struct {
	web.Media

	Title string `json:"title"`

	Fields  []Field  `json:"fields"`
	Records []Record `json:"records"`
}

// TableName table name
func (Form) TableName() string {
	return "survey_forms"
}

// Field field
type Field struct {
	web.Media

	Name      string `json:"name"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	SortOrder int    `json:"sortOrder"`

	FormID uint `json:"formId"`
	Form   Form
}

// TableName table name
func (Field) TableName() string {
	return "survey_fields"
}

// Record record
type Record struct {
	web.Model

	Value string `json:"value"`

	FormID uint `json:"formId"`
	Form   Form
}

// TableName table name
func (Record) TableName() string {
	return "survey_records"
}
