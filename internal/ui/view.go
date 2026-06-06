package ui

import (
	"html/template"
	"io"
)

type Data map[string]any

type View struct {
	TemplateName string
	Data         Data
}

func (v View) Render(w io.Writer, t *template.Template) error {
	return t.ExecuteTemplate(w, v.TemplateName, v.Data)
}
