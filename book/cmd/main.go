package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates {
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

type Contact struct {
    Name string
    Email string
}

func newContact(name string, email string) Contact {
    return Contact {
        Name: name,
        Email: email,
    }
}

type Contacts = []Contact

type Data struct {
    Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
    for _, contact := range d.Contacts {
        if contact.Email == email {
            return true
        }
    }
    return false
}

func newData() Data {
    return Data {
        Contacts: []Contact {
            newContact("Gavin", "hey@gav.codes"),
            newContact("Prime", "prime@gmail.com"),
        },
    }
}

type FormData struct {
    Values map[string]string
    Errors map[string]string
}

func newFormData() FormData {
    return FormData {
        Values: make(map[string]string),
        Errors: make(map[string]string),
    }
}

type Page struct {
    Data Data
    Form FormData
}

func newPage() Page {
    return Page {
        Data: newData(),
        Form: newFormData(),
    }
}

func main() {

    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/css", "css")

    page := newPage()
    e.Renderer = newTemplate()

    e.GET("/", func(c echo.Context) error {
        fmt.Println(c.RealIP(), "\n")
        return c.Render(200, "index", page)
    })

    e.POST("/contacts", func(c echo.Context) error {
        

        name := c.FormValue("name")
        email := c.FormValue("email")

        if page.Data.hasEmail(email) {
            formData := newFormData()
            formData.Values["name"] = name
            formData.Values["email"] = email

            formData.Errors["email"] = "Email already exists"

            c.Response().Header().Add("Hx-Redirect", "/404")
            return c.Render(422, "form", formData)
        }

        contact := newContact(name, email);
        page.Data.Contacts = append(page.Data.Contacts, contact)

        c.Render(200, "form", newFormData())
        return c.Render(200, "oob-contact", contact)
    })

    e.Logger.Fatal(e.Start(":420"))
}
