package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/jackc/pgx/v5"
)

func HandlePut(c echo.Context, conn *pgx.Conn) error {

    path := c.Param("path")

    switch path {

    }

    return nil
}
