package handlers

import (
	"gavrh/book/handlers/getHandlers"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandleGet(c echo.Context, jwtSecret string, pool *pgxpool.Pool) error {

    path := c.Param("path")

    switch path {
        case "login": return getHandlers.HandleGetLogin(c, jwtSecret, pool)
        case "signup": return getHandlers.HandleGetSignup(c, jwtSecret, pool)
        case "2fa": return getHandlers.HandleGetTwoFactor(c, jwtSecret, pool)

        // temp while has no favicon.ico
        case "favicon.ico":
            return nil

    }
    
    return getHandlers.HandleGetDefault(c, jwtSecret, pool)
}
