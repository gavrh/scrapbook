package main

import (
	"fmt"
	"gavrh/book/handlers"
	"gavrh/book/templates"

	// "io"
	"context"
	"os"

    "github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

    env, _ := godotenv.Read(".env")
    jwtSecret := env["JWT_SECRET"]
    pool, err := pgxpool.New(context.Background(), env["DATABASE_URL"])
    if err != nil {
        fmt.Println("could not connect to psql database.")
        os.Exit(1)
    }
    defer pool.Close()

    e := echo.New()
    e.Use(middleware.Logger())
    e.IPExtractor = echo.ExtractIPDirect()
    e.Static("/static/css", "css")
    e.Renderer = templates.NewTemplate()
    handlers.HandleRequests(e, jwtSecret, pool)
    // e.Logger.SetOutput(io.Discard) | disables logger
    e.Start(":420")

}
