package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
    return &Templates {
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

const (
    Index = "index"

    Login = "login"
    LoginForm = "login-form"
    Signup = "login"
    SignupForm = "login-form"
    TwoFactor = "2fa"
)
