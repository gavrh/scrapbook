package handlers

import (
    "gavrh/book/handlers/getHandlers"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HandleGet(c echo.Context, jwtSecret string, conn *pgx.Conn) error {

    path := c.Param("path")

    switch path {
        case "": getHandlers.HandleGetDefault(c, jwtSecret, conn)
        case "login": getHandlers.HandleGetLogin(c, jwtSecret, conn)
        case "signup": getHandlers.HandleGetSignup(c, jwtSecret, conn)
        case "2fa": getHandlers.HandleGetTwoFactor(c, jwtSecret, conn)

        // temp while has no favicon.ico
        case "favicon.ico":
            return nil

    }
    
    return c.Redirect(303, "/")
}
