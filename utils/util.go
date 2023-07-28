package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"html/template"
	"net/http"
	"p3/.gen/p3/public/model"
	. "p3/.gen/p3/public/table"
	"p3/db"
	"p3/entities"
	cErr "p3/err"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

func TokenVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Set("Email", "ankit@gmail.com")
		return next(c)
		SocialClientKey := "261940944085-n7a3smonr755otm67936q0rl169njkir.apps.googleusercontent.com"
		reqToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			return GetError(cErr.InvalidToken)
		}

		token := strings.TrimSpace(splitToken[1])

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
func AddRoles(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		email := c.Get("Email").(string)
		roles := GetRoles(email)
		c.Set("Roles", roles)
		return next(c)
	}
}

type EmployeeRoles struct {
	OrgId    *int64
	OrgName  string
	Role     string
	Managing map[string]bool
}
type Subordinates struct {
	model.Employees
	Manages []model.Employees `alias:"Manages"`
}

func GetRoles(email string) EmployeeRoles {
	s := GetEmployeesManaging(email)
	r := EmployeeRoles{
		Role:    s.Role,
		OrgId:   s.OrgID,
		OrgName: "test",
	}
	managing := map[string]bool{}
	for _, emp := range s.Manages {

		managing[emp.Email] = true
	}
	r.Managing = managing
	return r
}
func GetEmployeesManaging(email string) Subordinates {
	Manages := Employees.AS("Manages")
	var dest Subordinates
	stmt := SELECT(Employees.AllColumns, Manages.AllColumns).FROM(Employees.LEFT_JOIN(Manages, Employees.Email.EQ(Manages.ManagedBy))).WHERE(Employees.Email.EQ(String(email)))
	debugSql := stmt.DebugSql()
	fmt.Println(debugSql)
	database := db.GetJetDB()

	stmt.Query(database, &dest)
	fmt.Println(dest)
	return dest
}
