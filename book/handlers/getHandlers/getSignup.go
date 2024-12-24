package getHandlers

import (
	"gavrh/book/handlers/otherHandlers"
	"gavrh/book/templates"

	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandleGetSignup(c echo.Context, jwtSecret string, pool *pgxpool.Pool) error {
    if tokenCookie, tokenError := c.Cookie("token"); tokenError == nil {
        if _, twoFactorComplete, ok := otherHandlers.ValidateToken(tokenCookie, c.Request().RemoteAddr, jwtSecret); ok {
            if !twoFactorComplete {
                return c.Redirect(http.StatusSeeOther, "/2fa")
            }
            return c.Redirect(http.StatusSeeOther, "/")
        }
    }

    invite := c.QueryParam("invite")
    data := templates.NewLoginTemplate(false, "", "", invite)
    return c.Render(http.StatusOK, templates.Signup, data)
}
