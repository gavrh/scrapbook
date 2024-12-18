package main

import (
	"gavrh/cdn/encryption"

	"encoding/json"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
    "net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Error struct {
    Message string `json:"message"`
}
func Err(c echo.Context, code int, message string) error {
    err, _ := json.Marshal(Error { Message: message, })
    return c.JSON(code, err)
}

func interruptExit(channel chan os.Signal, key []byte) {
    <-channel
    os.Create(".env")
    f, _ := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    f.Write([]byte("ENCRYPTION_KEY="))
    f.Write(key)
    f.Close()
    os.Exit(0)
}

func main() {

    env, _ := godotenv.Read(".env")
    key := []byte(env["ENCRYPTION_KEY"])

    channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
    go interruptExit(channel, key)
    os.Remove(".env")

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
        os.Mkdir("storage/temp" + strconv.Itoa(randNumber), 0755)
        encryption.DecryptFile("storage/encrypted/" + file, "storage/temp" + strconv.Itoa(randNumber) + "/" + file + ".mkv", key)
        defer os.RemoveAll("storage/temp" + strconv.Itoa(randNumber))

        return c.File("storage/temp" + strconv.Itoa(randNumber) + "/" + file + ".mkv")
    })
    e.Start(":42069")

}
