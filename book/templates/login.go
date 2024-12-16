package templates


type LoginTemplate struct {
    IsLogin bool
    Username string
    Password string
    InviteCode string
}

func NewLoginTemplate(isLogin bool, username string, password string, inviteCode string) LoginTemplate {
    return LoginTemplate {
        IsLogin: isLogin,
        Username: username,
        Password: password,
        InviteCode: inviteCode,
    }
}
