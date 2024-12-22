package handlers

import (
    "gavrh/book/handlers/postHandlers"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandlePost(c echo.Context, jwtSecret string, pool *pgxpool.Pool) error {

    path := c.Param("path")

    switch path {

        case "login": return postHandlers.HandlePostLogin(c, jwtSecret, pool)
        case "signup": return nil
        case "2fa": return postHandlers.HandlePostTwoFactor(c, jwtSecret, pool)
        
    }

    return nil
}
