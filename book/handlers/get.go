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

func HandleGet(c echo.Context, conn *pgx.Conn) error {

    path := c.Param("path")

    fmt.Println(path)

    switch path {
        // "/"
        case "":
        // look for token cookie
        token, err := c.Cookie("token")
        if err != nil {
            data := templates.NewLoginTemplate(true, "", "", "")
            return c.Render(http.StatusOK, templates.Login, data)
        }

        var account_id string
        err = conn.QueryRow(context.Background(), "select account_id from accounts").Scan(&account_id)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        fmt.Println(account_id)
        fmt.Println(token)

        return c.Render(http.StatusOK, templates.Index, IndexData{ Id: account_id })

        // "/login"
        case "login":

        // "signup"
        case "signup":

    }
    
    return c.Redirect(303, "/")
}
