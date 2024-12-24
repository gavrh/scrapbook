package getHandlers

import (
	"gavrh/book/handlers/otherHandlers"
	"gavrh/book/templates"

	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandleGetLogin(c echo.Context, jwtSecret string, pool *pgxpool.Pool) error {
    if tokenCookie, tokenError := c.Cookie("token"); tokenError == nil {
        fmt.Println(tokenCookie.Value)
        if _, twoFactorComplete, ok := otherHandlers.ValidateToken(tokenCookie, c.Request().RemoteAddr, jwtSecret); ok {
            if !twoFactorComplete {
                return c.Redirect(http.StatusSeeOther, "/2fa")
            }
            return c.Redirect(http.StatusSeeOther, "/")
        }
    }

    data := templates.NewLoginTemplate(true, "", "", "")
    return c.Render(http.StatusOK, templates.Login, data)
}
