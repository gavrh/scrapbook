package templates

type TwoFactorTemplate struct {
    AccountId string
    Login string
    TwoFactorSecret string
    TwoFactorIsSetup bool
}

func NewTwoFactorTemplate(accountId string, login string, twoFactorSecret string, twoFactorIsSetup bool) TwoFactorTemplate {
    return TwoFactorTemplate {
        AccountId: accountId,
        Login: login,
        TwoFactorSecret: twoFactorSecret,
        TwoFactorIsSetup: twoFactorIsSetup,
    }
}
