package handlers

import (
	"gavrh/book/handlers/otherHandlers"
	"gavrh/book/templates"

	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HandlePost(c echo.Context, jwtSecret string, conn *pgx.Conn) error {

    path := c.Param("path")

    switch path {

    // "/login"
    case "login":

        login := c.FormValue("username")
        password := c.FormValue("password")

        var account_id string
        var account_2fa_secret string
        var account_setup_complete bool
        var user_login string
        err := conn.QueryRow(context.Background(),
            "SELECT account_id, account_2fa_secret, account_setup_complete, user_login FROM users " +
            "INNER JOIN accounts USING(account_id)" +
            "WHERE user_login = '" + login + "' " +
            "AND user_password = '" + password + "'",
        ).Scan(&account_id, &account_2fa_secret, &account_setup_complete, &user_login)
        if err != nil {
            fmt.Println(err)
            data := templates.NewLoginTemplate(true, login, "", "")
            return c.Render(http.StatusOK, templates.Login, data)
        }
        
        token, err := otherHandlers.CreateToken(account_id, c.Request().RemoteAddr, false, jwtSecret)

        c.Response().Header().Add("Hx-Push-Url", "/2fa")
        c.Response().Header().Add("Set-Cookie", "token="+token+"; domain=localhost;")

        data := templates.NewTwoFactorTemplate(account_id, user_login, account_2fa_secret, account_setup_complete)
        return c.Render(http.StatusOK, templates.TwoFactor, data)
        

    // "/signup"
    case "signup":
        
    }

    return nil
}
