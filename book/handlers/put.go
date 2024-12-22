package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HandlePut(c echo.Context, pool *pgxpool.Pool) error {

    path := c.Param("path")

    switch path {

    }

    return nil
}
