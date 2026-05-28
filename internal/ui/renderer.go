package ui

import (
	"html/template"
	"io"
)

type Data map[string]any

type Renderer interface {
	Render(w io.Writer, t *template.Template) error
}

type Group []Renderer

func (g Group) Render(w io.Writer, t *template.Template) error {
	for _, item := range g {
		if err := item.Render(w, t); err != nil {
			return err
		}
	}
	return nil
}

func Compose(items ...Renderer) Group {
	return Group(items)
}
