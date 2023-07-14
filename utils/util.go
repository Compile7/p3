package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	"p3/entities"
	cErr "p3/err"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

func TokenVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		SocialClientKey := "327238763639-da8uschea6n56fvnmu47o8ve9r65eee3.apps.googleusercontent.com"
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return GetError(cErr.TokenMissing)
			//return c.JSON(http.StatusForbidden, "fxdbfgsb")
		}
		payload, err := idtoken.Validate(context.Background(), token, SocialClientKey)
		if err != nil {
			return GetError(cErr.InvalidToken)
		}
		var SocialClaims entities.SocialTokenClaims
		mapErr := mapstructure.Decode(payload.Claims, &SocialClaims)
		if mapErr != nil {
			return GetError(cErr.Mapping)
		}
		fmt.Println(payload)

		now := time.Now().UTC().Unix()

		fmt.Println(SocialClaims.ExpiryTime)

		if !SocialClaims.IsNotExpired(now) {

			//handle error
			return echo.ErrBadRequest
		}
		if !SocialClaims.IsIssueTimeValid(now) {
			//handle error
			return echo.ErrBadRequest
		}

		if !SocialClaims.IsAudienceValid(SocialClientKey) {
			//handle error
			return echo.ErrBadRequest
		}

		if !SocialClaims.IsIssuerValid() {
			//handle error
			return echo.ErrBadRequest
		}

		if !SocialClaims.DoesApiKeyMatch(SocialClientKey) {
			//handle error
			return echo.ErrBadRequest
		}
		if SocialClaims.EmailNotVerified() {
			//handle error
			return echo.ErrBadRequest
		}
		c.Set("Email", SocialClaims.Email)
		c.Set("Name", SocialClaims.FullName)
		c.Set("ProfilePicture", SocialClaims.ProfilePicture)
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

	Name string `json:"name"`
}
type BrevoData struct {
	Sender      NameEmail   `json:"sender"`
	To          []NameEmail `json:"to"`
	Subject     string      `json:"subject"`
	HtmlContent string      `json:"htmlContent"`
}

func InviteUser(invite InviteEmployee, emailOptions EmailOptions) {
	client := &http.Client{}
	tmpl, e := template.ParseFiles("utils/email.html")
	fmt.Println(e)
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, invite)
	apikey, _ := viper.Get("brevoApiKey").(string)

	body := tpl.String()
	brevo := BrevoData{
		Sender:      NameEmail{Email: "ankit.aabad1@gmail.com", Name: "ankit"},
		To:          []NameEmail{{Email: emailOptions.Receiver, Name: "add receiver name here"}},
		Subject:     "Invitation to Join P3",
		HtmlContent: body,
	}
	brevoOutput, _ := json.Marshal(brevo)
	req, _ := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewReader(brevoOutput))
	req.Header.Set("api-key", apikey)

	res, e := client.Do(req)
	fmt.Println(res)
}

func GetError(customError cErr.CustomError) error {
	return echo.NewHTTPError(customError.Status, customError.Message)
}
