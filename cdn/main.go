package main

import (
	"gavrh/cdn/encryption"

	"encoding/json"
	"os"
	"os/signal"
	"syscall"
    "math/rand"
    "strconv"

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

    // encryption.EncryptFile("storage/users/picture.jpeg", "storage/encrypted/picture", key)
    

    e := echo.New()
    e.Use(middleware.Logger())
    // e.Logger.SetOutput(io.Discard) | disables logger
    e.GET("/:path/:file", func(c echo.Context) error {

        // path := c.Param("path")
        file := c.Param("file")

        f, err := os.Open("storage/encrypted/" + file)
        if err != nil {
            Err(c, 404, "Not Found")
        }
        defer f.Close()

        // token authorization check
        _, err = c.Cookie("token")
        if err != nil {
            Err(c, 401, "Not Authorized")
        }

        randNumber := rand.Intn(1000) + 1
        os.Mkdir("storage/temp" + strconv.Itoa(randNumber), 0755)
        encryption.DecryptFile("storage/encrypted/" + file, "storage/temp" + strconv.Itoa(randNumber) + "/picture.jpeg", key)
        defer os.RemoveAll("storage/temp" + strconv.Itoa(randNumber))

        return c.File("storage/temp" + strconv.Itoa(randNumber) + "/" + file + ".jpeg")
    })
    e.Start(":42069")

}
