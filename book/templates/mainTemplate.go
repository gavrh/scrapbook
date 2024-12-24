package templates

type MainTemplate struct {
    Name string
}

func NewMainTemplate(name string) MainTemplate {
    return MainTemplate {
        Name: name,
    }
}
