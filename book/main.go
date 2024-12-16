package main

import (
	"gavrh/book/handlers"
	"gavrh/book/templates"

	// "io"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    "github.com/joho/godotenv"
)

func main() {

    env, _ := godotenv.Read(".env")
    conn, err := pgx.Connect(context.Background(), env["DATABASE_URL"])
    if err != nil {
        os.Exit(1)
    }
    defer conn.Close(context.Background())

    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/static/css", "css")
    e.Renderer = templates.NewTemplate()
    handlers.HandleRequests(e, conn)
    // e.Logger.SetOutput(io.Discard) | disables logger
    e.Start(":420")

}
