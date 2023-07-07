package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

func TokenVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		payload, err := idtoken.Validate(context.Background(), token, "711686112966-qme1urtstdo0thlv10v0rjv62nelpf5r.apps.googleusercontent.com")
		if err != nil {

		}
		fmt.Println(payload)
		return next(c)
	}
}

type InviteEmployee struct {
	Link         string
	Organisation string
}
type EmailOptions struct {
	Sender   string
	Receiver string
	Body     string
}
type NameEmail struct {
	Email string `json:"email"`

	Name  string `json:"name"`
}
type BrevoData struct {
	Sender []NameEmail `json:"sender"`
	To     []NameEmail `json:"to"`
	Subject string `json:"subject"`
	HtmlContent string `json:"htmlContent"`
}

func InviteUser(invite InviteEmployee, emailOptions EmailOptions) {
	client := &http.Client{}
	tmpl, _ := template.ParseFiles("email.html")
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, invite)

	body := tpl.String()
	brevo := BrevoData{
		Sender: []NameEmail{{Email:emailOptions.Sender}},
		To: []NameEmail{{Email:emailOptions.Receiver}},
		Subject : "Invitation to Join P3",
		HtmlContent: body,
	}
	brevoOutput, _ := json.Marshal(brevo)
	req, _ := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewReader(brevoOutput))
	client.Do(req)
}
