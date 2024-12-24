package handlers

import (
	"gavrh/book/templates"
	"gavrh/book/handlers/getHandlers"

	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandleRequests(e *echo.Echo, jwtSecret string, pool *pgxpool.Pool) {

    // get
    e.GET("/", func (c echo.Context) error { return getHandlers.HandleGetDefault(c, jwtSecret, pool) })
    e.GET("/:path", func (c echo.Context) error { return HandleGet(c, jwtSecret, pool) })

    // post
    e.POST("/:path", func (c echo.Context) error { return HandlePost(c, jwtSecret, pool) })

    // put
    e.PUT("/:path", func (c echo.Context) error { return HandlePut(c, pool) })

    e.PUT("/login", func (c echo.Context) error {
        data := templates.NewLoginTemplate(true, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.LoginForm, data)
    })
    e.PUT("/signup", func (c echo.Context) error {
        data := templates.NewLoginTemplate(false, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.SignupForm , data)
    })
}
