package main

import (
    "gavrh/book/handlers"
    "gavrh/book/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/static/css", "css")
    e.Renderer = templates.NewTemplate()
    handlers.HandleRequests(e)
    e.Logger.Fatal(e.Start(":420"))

}
