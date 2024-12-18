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

func TokenExists(c echo.Context) (*http.Cookie, error) {
    token, err := c.Cookie("token")
    return token, err
}

func HandleRequests(e *echo.Echo, conn *pgx.Conn) {

    // get
    e.GET("/", func (c echo.Context) error { return HandleGet(c, conn) })
    e.GET("/:path", func (c echo.Context) error { return HandleGet(c, conn) })

    // post

    // put
    e.PUT("/:path", func (c echo.Context) error { return HandlePut(c, conn) })

    e.GET("/login", func(c echo.Context) error {

        token, _ := c.Cookie("token")
        if token != nil {
            var account_id string
            err := conn.QueryRow(context.Background(), "select account_id from accounts").Scan(&account_id)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            fmt.Println(account_id)
            fmt.Println(token)
            return c.Render(http.StatusOK, templates.Index, IndexData{ Id: account_id })
        }

        data := templates.NewLoginTemplate(true, "", "", "")
        return c.Render(http.StatusOK, templates.Login, data)
    })
    e.PUT("/login", func (c echo.Context) error {
        data := templates.NewLoginTemplate(true, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.LoginForm, data)
    })
    e.GET("/signup", func(c echo.Context) error {
        token, _ := c.Cookie("token")
        if token != nil {
            var account_id string
            err := conn.QueryRow(context.Background(), "select account_id from accounts").Scan(&account_id)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            fmt.Println(account_id)
            fmt.Println(token)
            return c.Render(http.StatusOK, templates.Index, IndexData{ Id: account_id })
        }
        invite := c.QueryParam("invite")
        data := templates.NewLoginTemplate(false, "", "", invite)
        return c.Render(http.StatusOK, templates.Signup, data)
    })
    e.PUT("/signup", func (c echo.Context) error {
        data := templates.NewLoginTemplate(false, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.SignupForm , data)
    })
}
