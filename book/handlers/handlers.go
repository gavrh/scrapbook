package handlers

import (
    "gavrh/book/templates"

    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/jackc/pgx/v5"
)

func HandleRequests(e *echo.Echo, jwtSecret string, conn *pgx.Conn) {

    // get
    e.GET("/", func (c echo.Context) error { return HandleGet(c, jwtSecret, conn) })
    e.GET("/:path", func (c echo.Context) error { return HandleGet(c, jwtSecret, conn) })

    // post
    e.POST("/:path", func (c echo.Context) error { return HandlePost(c, jwtSecret, conn) })

    // put
    e.PUT("/:path", func (c echo.Context) error { return HandlePut(c, conn) })

    e.PUT("/login", func (c echo.Context) error {
        data := templates.NewLoginTemplate(true, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.LoginForm, data)
    })
    e.PUT("/signup", func (c echo.Context) error {
        data := templates.NewLoginTemplate(false, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.SignupForm , data)
    })
}
