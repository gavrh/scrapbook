package main

import (
	"fmt"
	"gavrh/cdn/encryption"
	"strings"

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
    site := env["SITE_URL"]

    // encryption.EncryptFile("storage/users/test.mkv", "storage/encrypted/test", key)

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
        AllowOrigins: []string{ site },
        AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
        AllowMethods: []string{ http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete },
    }))

    e.GET("/:path/:file", func(c echo.Context) error {

        // path := c.Param("path")
        fileParam := c.Param("file")
        fileParts := strings.Split(fileParam, ".")
        file := fileParts[0]
        var ext string
        if len(fileParts) > 1 {
            ext = "." + fileParts[1]
        }


        f, err := os.Open("storage/encrypted/" + file)
        if err != nil {
            Err(c, 404, "Not Found")
        }
        defer f.Close()

        randNumber := rand.Intn(1000) + 1
        os.Mkdir("storage/.temp" + strconv.Itoa(randNumber), 0755)
        encryption.DecryptFile("storage/encrypted/" + file, "storage/.temp" + strconv.Itoa(randNumber) + "/" + file + ext, key)
        defer os.RemoveAll("storage/.temp" + strconv.Itoa(randNumber))

        return c.File("storage/.temp" + strconv.Itoa(randNumber) + "/" + file + ext)
    })

    e.GET("/qr", func (c echo.Context) error {
        
        data := c.QueryParam("data")
        id := c.QueryParam("id")
        if len(data) == 0 || len(id) == 0 {
            return nil
        }

        err := qrcode.WriteFile(data, qrcode.Medium, 240, "storage/qr-" + id + ".png")
        if err != nil {
            fmt.Println(err)
            return nil
        }
        defer os.Remove("storage/qr-" + id + ".png")

        return c.File("storage/qr-" + id + ".png")
    })

    e.Start(":42069")

}
