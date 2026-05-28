package ui

import (
	"html/template"
	"net/http"
)

const TurboStreamContentType = "text/vnd.turbo-stream.html"

type Engine struct {
	tmpl *template.Template
}

func NewEngine(tmpl *template.Template) *Engine {
	return &Engine{tmpl: tmpl}
}

func (e *Engine) WriteHTML(w http.ResponseWriter, view Renderer) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return view.Render(w, e.tmpl)
}

func (e *Engine) WriteTurbo(w http.ResponseWriter, view Renderer) error {
	w.Header().Set("Content-Type", TurboStreamContentType)
	return view.Render(w, e.tmpl)
}
