package getHandlers

import (
	"gavrh/book/handlers/otherHandlers"
	"gavrh/book/templates"

	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HandleGetLogin(c echo.Context, jwtSecret string, conn *pgx.Conn) error {
    tokenCookie, tokenError := c.Cookie("token")
    if tokenError == nil {
        account_id, twoFactorComplete, ok := otherHandlers.ValidateToken(tokenCookie, tokenError, c.Request().RemoteAddr, jwtSecret)
        if ok {
            var account_2fa_secret string
            var account_setup_complete bool
            var user_login string
            err := conn.QueryRow(context.Background(),
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

            return c.Render(http.StatusOK, templates.Index, otherHandlers.IndexData { Id: account_id })
        }
    }

    data := templates.NewLoginTemplate(true, "", "", "")
    return c.Render(http.StatusOK, templates.Login, data)
}
