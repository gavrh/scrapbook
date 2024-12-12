package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates {
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

type LoginData struct {
    LoginType string
    OtherLoginType string
}

func main() {

    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/css", "css")
    e.Renderer = newTemplate()

    e.GET("/", func(c echo.Context) error {

        // look for token cookie
        cookies := c.Cookies()
        var token string
        for _, cookie := range cookies {
            if cookie.Name == "token" {
                token = cookie.Value
                break
            }
        }

        fmt.Println(token)

        if len(token) != 0 {
            return c.Render(http.StatusOK, "index", nil)
        }

        fmt.Println("Should redirect")

        return c.Render(http.StatusOK, "login", LoginData{ LoginType: "login", OtherLoginType: "signup" })
    })

    e.GET("/login", func(c echo.Context) error {
        return c.Render(http.StatusOK, "login", LoginData{ LoginType: "login", OtherLoginType: "signup" })
    })
    e.GET("/signup", func(c echo.Context) error {
        return c.Render(http.StatusOK, "login", LoginData{ LoginType: "signup", OtherLoginType: "login" })
    })

    e.Logger.Fatal(e.Start(":420"))
}
