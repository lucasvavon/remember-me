package views

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

// Implement e.Renderer interface
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
