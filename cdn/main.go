package main

import (
	"os"
    "encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Error struct {
    Message string `json:"message"`
}
func NewError(message string) Error {
    return Error {
        Message: message,
    }
}

func main() {

    e := echo.New()
    e.Use(middleware.Logger())
    // e.Logger.SetOutput(io.Discard) | disables logger
    e.GET("/:path/:file", func(c echo.Context) error {

        path := c.Param("path")
        file := c.Param("file")

        f, err := os.Open("storage/" + path + "/" + file)
        if err != nil {
            err, _ := json.Marshal(NewError("File Not Found"))
            return c.JSON(404, string(err))
        }
        defer f.Close()

        _, err = c.Cookie("token")
        if err != nil {
            err, _ := json.Marshal(NewError("Not Authorized"))
            return c.JSON(401, string(err))
        }

        return c.Stream(200, "image/png", f)
    })
    e.Start(":42069")

}
