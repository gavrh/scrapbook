package handlers

import (
    "gavrh/book/templates"

    "fmt"
    "net/http"
    "context"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/jackc/pgx/v5"
)

type IndexData struct {
    Id string
}

func HandleRequests(e *echo.Echo, conn *pgx.Conn) {
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

        if len(token) != 0 {

            var account_id string
            err := conn.QueryRow(context.Background(), "select account_id from accounts").Scan(&account_id)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }

            fmt.Println(account_id)
            fmt.Println("test")

            return c.Render(http.StatusOK, templates.Index, IndexData{ Id: account_id })
        }

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
