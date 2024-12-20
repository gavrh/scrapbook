package handlers

import (
    "gavrh/book/templates"

    "fmt"
    "net/http"
    "context"

    "github.com/labstack/echo/v4"
    "github.com/jackc/pgx/v5"
)

type IndexData struct {
    Id string
}

func HandleRequests(e *echo.Echo, jwtSecret string, conn *pgx.Conn) {

    // get
    e.GET("/", func (c echo.Context) error { return HandleGet(c, jwtSecret, conn) })
    e.GET("/:path", func (c echo.Context) error { return HandleGet(c, jwtSecret, conn) })

    // post
    e.POST("/:path", func (c echo.Context) error { return HandlePost(c, jwtSecret, conn) })

    // put
    e.PUT("/:path", func (c echo.Context) error { return HandlePut(c, conn) })

    e.GET("/login", func(c echo.Context) error {
        // look for token cookie
        token, err := c.Cookie("token")
        if err == nil {
            account_id, twoFactorComplete, ok := ValidateToken(token.Value, c.Request().RemoteAddr, jwtSecret)
            if ok {
                var account_2fa_secret string
                var account_setup_complete bool
                var user_login string
                err = conn.QueryRow(context.Background(),
                    "SELECT account_id, account_2fa_secret, account_setup_complete, user_login FROM users " +
                    "INNER JOIN accounts USING(account_id)" +
                    "WHERE account_id = '" + account_id + "'",
                    ).Scan(&account_id, &account_2fa_secret, &account_setup_complete, &user_login)
                if err != nil {
                    fmt.Println(err)
                }

                if !twoFactorComplete {
                    data := templates.NewTwoFactorTemplate(account_id, user_login, account_2fa_secret, account_setup_complete)
                    return c.Render(http.StatusOK, templates.TwoFactor, data)
                }
                
                return c.Render(http.StatusOK, templates.Index, IndexData { Id: account_id })
            }
        }

        data := templates.NewLoginTemplate(true, "", "", "")
        return c.Render(http.StatusOK, templates.Login, data)
    })
    e.PUT("/login", func (c echo.Context) error {
        data := templates.NewLoginTemplate(true, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.LoginForm, data)
    })
    e.GET("/signup", func(c echo.Context) error {
        // look for token cookie
        token, err := c.Cookie("token")
        if err == nil {
            account_id, twoFactorComplete, ok := ValidateToken(token.Value, c.Request().RemoteAddr, jwtSecret)
            if ok {
                var account_2fa_secret string
                var account_setup_complete bool
                var user_login string
                err = conn.QueryRow(context.Background(),
                    "SELECT account_id, account_2fa_secret, account_setup_complete, user_login FROM users " +
                    "INNER JOIN accounts USING(account_id)" +
                    "WHERE account_id = '" + account_id + "'",
                    ).Scan(&account_id, &account_2fa_secret, &account_setup_complete, &user_login)
                if err != nil {
                    fmt.Println(err)
                }

                if !twoFactorComplete {
                    data := templates.NewTwoFactorTemplate(account_id, user_login, account_2fa_secret, account_setup_complete)
                    return c.Render(http.StatusOK, templates.TwoFactor, data)
                }
                
                return c.Render(http.StatusOK, templates.Index, IndexData { Id: account_id })
            }
        }

        invite := c.QueryParam("invite")
        data := templates.NewLoginTemplate(false, "", "", invite)
        return c.Render(http.StatusOK, templates.Signup, data)
    })
    e.PUT("/signup", func (c echo.Context) error {
        data := templates.NewLoginTemplate(false, c.FormValue("username"), c.FormValue("password"), "")
        return c.Render(http.StatusOK, templates.SignupForm , data)
    })
}
