package otherHandlers

import (
	"fmt"
	"time"

	"github.com/xlzd/gotp"
)

func VerifyOTP(code string, twoFactorSecret string) error {
    totp := gotp.NewDefaultTOTP(twoFactorSecret)
    if totp.Verify(code, time.Now().Unix()) {
        return nil
    }
    return fmt.Errorf("Incorrect value")
}
