package postHandlers

import (
	"fmt"
	"gavrh/book/handlers/otherHandlers"
	"gavrh/book/templates"

	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandlePostLogin(c echo.Context, jwtSecret string, pool *pgxpool.Pool) error {
    login := c.FormValue("username")
    password := c.FormValue("password")

    var account_id string
    var account_2fa_secret string
    var account_setup_complete bool
    var user_login string
    err := pool.QueryRow(context.Background(),
        "SELECT account_id, account_2fa_secret, account_setup_complete, user_login FROM users " +
        "INNER JOIN accounts USING(account_id)" +
        "WHERE user_login = '" + login + "' " +
        "AND user_password = '" + password + "'" +
        "LIMIT 1;",
        ).Scan(&account_id, &account_2fa_secret, &account_setup_complete, &user_login)
    if err != nil {
        fmt.Println(err)
        data := templates.NewLoginTemplate(true, login, "", "")
        return c.Render(http.StatusOK, templates.Login, data)
    }

    token, err := otherHandlers.CreateToken(account_id, c.RealIP(), false, jwtSecret)

    c.Response().Header().Add("Hx-Replace-Url", "/2fa")
    c.Response().Header().Add("Set-Cookie", "token="+token+"; domain=localhost;")
    data := templates.NewTwoFactorTemplate(account_id, user_login, account_2fa_secret, account_setup_complete)
    return c.Render(http.StatusOK, templates.TwoFactor, data)
}
