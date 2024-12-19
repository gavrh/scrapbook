package main

import (
	"fmt"
	"gavrh/cdn/encryption"

	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/skip2/go-qrcode"
)

type Error struct {
    Message string `json:"message"`
}
func Err(c echo.Context, code int, message string) error {
    err, _ := json.Marshal(Error { Message: message, })
    return c.JSON(code, err)
}

func main() {

    env, _ := godotenv.Read(".env")
    key := []byte(env["ENCRYPTION_KEY"])

    // encryption.EncryptFile("storage/users/prime.mkv", "storage/encrypted/prime", key)

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
        AllowOrigins: []string{ "https://scrapbook.sytes.net" },
        AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
        AllowMethods: []string{ http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete },
    }))

    e.GET("/:path/:file", func(c echo.Context) error {

        // path := c.Param("path")
        file := c.Param("file")

        f, err := os.Open("storage/encrypted/" + file)
        if err != nil {
            Err(c, 404, "Not Found")
        }
        defer f.Close()

        randNumber := rand.Intn(1000) + 1
        os.Mkdir("storage/.temp" + strconv.Itoa(randNumber), 0755)
        encryption.DecryptFile("storage/encrypted/" + file, "storage/.temp" + strconv.Itoa(randNumber) + "/" + file + ".mkv", key)
        defer os.RemoveAll("storage/.temp" + strconv.Itoa(randNumber))

        return c.File("storage/.temp" + strconv.Itoa(randNumber) + "/" + file + ".mkv")
    })

    e.GET("/qr", func (c echo.Context) error {
        
        data := c.QueryParam("data")
        id := c.QueryParam("id")
        if len(data) == 0 || len(id) == 0 {
            return nil
        }

        err := qrcode.WriteFile(data, qrcode.Medium, 256, "storage/qr-" + id + ".png")
        if err != nil {
            fmt.Println(err)
            return nil
        }
        defer os.Remove("storage/qr-" + id + ".png")

        return c.File("storage/qr-" + id + ".png")
    })

    e.Start(":42069")

}
