package handlers

import (
	"context"
	"gavrh/book/templates"

	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HandleGet(c echo.Context, jwtSecret string, conn *pgx.Conn) error {

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

        account_id, twoFactorComplete, ok := ValidateToken(token.Value, c.Request().RemoteAddr, jwtSecret)
        if !ok {
            data := templates.NewLoginTemplate(true, "", "", "")
            fmt.Println("ERROR VALIDATING TOKEN")
            return c.Render(http.StatusOK, templates.Login, data)
        }

        var account_2fa_secret string
        var account_setup_complete bool
        var user_login string
        err = conn.QueryRow(context.Background(),
            "SELECT account_id, account_2fa_secret, account_setup_complete, user_login FROM users " +
            "INNER JOIN accounts USING(account_id)" +
            "WHERE account_id = '" + account_id + "'",
        ).Scan(&account_id, &account_2fa_secret, &account_setup_complete, &user_login)
        if err != nil {
            fmt.Println(err)
        }

        if !twoFactorComplete {
            data := templates.NewTwoFactorTemplate(account_id, user_login, account_2fa_secret, account_setup_complete)
            return c.Render(http.StatusOK, templates.TwoFactor, data)
        }

        return c.Render(http.StatusOK, templates.Index, IndexData{ Id: account_id })

        // "/login"
        case "login":

        // "signup"
        case "signup":

        case "2fa":
        // look for token cookie
        token, err := c.Cookie("token")
        if err != nil {
            data := templates.NewLoginTemplate(true, "", "", "")
            return c.Render(http.StatusOK, templates.Login, data)
        }

        account_id, twoFactorComplete, ok := ValidateToken(token.Value, c.Request().RemoteAddr, jwtSecret)
        if !ok {
            data := templates.NewLoginTemplate(true, "", "", "")
            fmt.Println("ERROR VALIDATING TOKEN")
            return c.Render(http.StatusOK, templates.Login, data)
        }

        var account_2fa_secret string
        var account_setup_complete bool
        var user_login string
        err = conn.QueryRow(context.Background(),
            "SELECT account_id, account_2fa_secret, account_setup_complete, user_login FROM users " +
            "INNER JOIN accounts USING(account_id)" +
            "WHERE account_id = '" + account_id + "'",
        ).Scan(&account_id, &account_2fa_secret, &account_setup_complete, &user_login)
        if err != nil {
            fmt.Println(err)
        }

        if twoFactorComplete {
            return c.Render(http.StatusOK, templates.Index, IndexData { Id: account_id })
        }

        data := templates.NewTwoFactorTemplate(account_id, user_login, account_2fa_secret, account_setup_complete)
        return c.Render(http.StatusOK, templates.TwoFactor, data)
            

        // temp while has no favicon.ico
        case "favicon.ico":
            return nil

    }
    
    return c.Redirect(303, "/")
}
