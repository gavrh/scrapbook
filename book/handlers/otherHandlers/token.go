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

func CreateToken(accountId string, remoteIp string, twoFactor bool, jwtSecret string) (string , error) {

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
        "account_id": accountId,
        "remote_ip": remoteIp,
        "two_factor": twoFactor,
        "created_at": time.Now().UTC().UnixMilli(),
    })

    return token.SignedString([]byte(jwtSecret))
} 

func ValidateToken(tokenCookie *http.Cookie, tokenError error, remoteIp string, jwtSecret string) (string, bool,  bool) {
    tokenString := tokenCookie.Value
    var nilString string
    var nilBool bool
    if tokenError != nil || len(tokenString) == 0 || tokenString == "deleted" {
        return nilString, nilBool, false
    }

    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
        
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(jwtSecret), nil
    })
    if err != nil {
        return nilString, nilBool, false
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        if remoteIp == claims["remote_ip"].(string) {
            return claims["account_id"].(string), claims["two_factor"].(bool), true
        }
    }
    return nilString, nilBool, false
}
