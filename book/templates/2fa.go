package templates

type TwoFactorTemplate struct {
    AccountId string
    Login string
    TwoFactorSecret string
    AccountSetupComplete bool
}

func NewTwoFactorTemplate(accountId string, login string, twoFactorSecret string, accountSetupComplete bool) TwoFactorTemplate {
    return TwoFactorTemplate {
        AccountId: accountId,
        Login: login,
        TwoFactorSecret: twoFactorSecret,
        AccountSetupComplete: accountSetupComplete,
    }
}
