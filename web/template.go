package web

import (
	"io"
	"html/template"

	"github.com/labstack/echo/v4"
)

var _ echo.Renderer = (*TemplatesRenderer)(nil)

// Implement echo.Renderer interface
type TemplatesRenderer struct {
	templates *template.Template
}

func (t *TemplatesRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}