package ui

import (
	"html/template"
	"net/http"
)

const (
	TurboStreamContentType = "text/vnd.turbo-stream.html"
	HTMLContentType        = "text/html; charset=utf-8"
)

type Engine struct {
	tmpl *template.Template
}

func NewEngine(tmpl *template.Template) *Engine {
	return &Engine{tmpl: tmpl}
}

func (e *Engine) WriteHTML(w http.ResponseWriter, view View) error {
	w.Header().Set("Content-Type", HTMLContentType)
	return view.Render(w, e.tmpl)
}

func (e *Engine) WriteTurbo(w http.ResponseWriter, view View) error {
	w.Header().Set("Content-Type", TurboStreamContentType)
	return view.Render(w, e.tmpl)
}
