package main

import (
    "gavrh/book/templates"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoginData struct {
    LoginType string
    OtherLoginType string
}

func main() {

    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/static/css", "css")
    e.Renderer = templates.NewTemplate()

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

        data := templates.NewLoginTemplate(true, "", "", "")
        return c.Render(http.StatusOK, "login", data)
    })

    e.GET("/login", func(c echo.Context) error {
        data := templates.NewLoginTemplate(true, "", "", "")
        return c.Render(http.StatusOK, "login", data)
    })
    e.PUT("/login", func (c echo.Context) error {
        data := templates.NewLoginTemplate(true, "", "", "")
        c.Request().Header.Add("Hx-Replace-Url", "/login")
        return c.Render(http.StatusOK, "login-form", data)
    })
    e.GET("/signup", func(c echo.Context) error {
        invite := c.QueryParam("invite")
        data := templates.NewLoginTemplate(false, "", "", invite)
        return c.Render(http.StatusOK, "login", data)
    })
    e.PUT("/signup", func (c echo.Context) error {
        data := templates.NewLoginTemplate(false, "", "", "")
        c.Request().Header.Add("Hx-Replace-Url", "/login")
        return c.Render(http.StatusOK, "login-form", data)
    })

    e.Logger.Fatal(e.Start(":420"))
}
