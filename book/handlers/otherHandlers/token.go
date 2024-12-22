package otherHandlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type IndexData struct {
    Id string
}

func CreateToken(accountId string, ipAddr string, twoFactor bool, jwtSecret string) (string , error) {

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
        "account_id": accountId,
        "ip_addr": ipAddr,
        "two_factor": twoFactor,
        "created_at": time.Now().UTC().UnixMilli(),
    })

    return token.SignedString([]byte(jwtSecret))
} 

func ValidateToken(tokenCookie *http.Cookie, remoteIp string, jwtSecret string) (string, bool,  bool) {
    tokenString := tokenCookie.Value
    var nilString string
    var nilBool bool

    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            fmt.Println("HERE 1")
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(jwtSecret), nil
    })
    if err != nil {
        fmt.Println("HERE 2")
        return nilString, nilBool, false
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        fmt.Println(claims["ip_addr"])
        return claims["account_id"].(string), claims["two_factor"].(bool), true
    }
    fmt.Println("HERE 3")
    return nilString, nilBool, false
}
