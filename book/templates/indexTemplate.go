package templates

type IndexTemplate struct {
    AccountId string
    UserLogin string
    MainDisplay MainTemplate
}

func NewIndexTemplate(accountId string, userLogin string, mainDisplay MainTemplate) IndexTemplate {
    return IndexTemplate {
        AccountId: accountId,
        UserLogin: userLogin,
        MainDisplay: mainDisplay,
    }
}
