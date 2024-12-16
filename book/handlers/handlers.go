package handlers

import (
    "gavrh/book/templates"

    "fmt"
    "net/http"

    "github.com/labstack/echo/v4"
)

func HandleRequests(e *echo.Echo) {
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
            return c.Render(http.StatusOK, templates.Index, nil)
        }

        fmt.Println("Should redirect")

        data := templates.NewLoginTemplate(true, "", "", "")
        return c.Render(http.StatusOK, templates.Login, data)
    })

    e.GET("/login", func(c echo.Context) error {
        data := templates.NewLoginTemplate(true, "", "", "")
        return c.Render(http.StatusOK, templates.Login, data)
    })
    e.PUT("/login", func (c echo.Context) error {
        data := templates.NewLoginTemplate(true, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.LoginForm, data)
    })
    e.GET("/signup", func(c echo.Context) error {
        invite := c.QueryParam("invite")
        data := templates.NewLoginTemplate(false, "", "", invite)
        return c.Render(http.StatusOK, templates.Signup, data)
    })
    e.PUT("/signup", func (c echo.Context) error {
        data := templates.NewLoginTemplate(false, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.SignupForm , data)
    })
}
