package postHandlers

import (
	"gavrh/book/handlers/otherHandlers"
	"gavrh/book/templates"

	"fmt"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func HandlePostTwoFactor(c echo.Context, jwtSecret string, pool *pgxpool.Pool) error {
    
    code := c.FormValue("code")
    account_id := c.FormValue("account_id")
    // user_login := c.FormValue("user_login")
    account_2fa_secret := c.FormValue("account_2fa_secret")
    // account_setup_complete := c.FormValue("account_setup_complete")
    fmt.Println(c)
    fmt.Println(code)
    err := otherHandlers.VerifyOTP(code, account_2fa_secret)
    if err != nil {
        fmt.Println(err)
        return c.Render(http.StatusNotFound, "", nil)
    }

    _, err = pool.Query(context.Background(), "update accounts set account_setup_complete = true where account_id = '" + account_id + "';")
    if err != nil {
        fmt.Println(err)
    }

    token, err := otherHandlers.CreateToken(account_id, c.RealIP(), true, jwtSecret)

    c.Response().Header().Add("Hx-Replace-Url", "/")
    c.Response().Header().Add("Set-Cookie", "token="+token+"; domain=localhost;")
    main := templates.NewMainTemplate("HOME")
    return c.Render(http.StatusOK, templates.Index, templates.NewIndexTemplate(account_id, "gavin", main))
}
