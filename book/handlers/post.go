package handlers

import (
    "gavrh/book/handlers/postHandlers"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HandlePost(c echo.Context, jwtSecret string, conn *pgx.Conn) error {

    path := c.Param("path")

    switch path {

        case "login": return postHandlers.HandlePostLogin(c, jwtSecret, conn)
        case "signup": return nil
        case "2fa": return postHandlers.HandlePostTwoFactor(c, jwtSecret, conn)
        
    }

    return nil
}
