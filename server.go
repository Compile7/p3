package main

import (
	"fmt"
	"p3/db"
	handlers "p3/handler"
	"p3/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		var errString string
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				if e.Tag() == "notblank" {
					errString += fmt.Sprintf("The value of %s can not be null or empty.", e.Field())
				} else if e.Tag() == "email" {
					errString += fmt.Sprintf("The value of %s should be a valid email.", e.Field())
				} else {
					errString += fmt.Sprintf("The %s is a required parameter.", e.Field())
				}
			}
			return fmt.Errorf(errString)
		}
		return nil
	}
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(utils.TokenVerify)
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	dbInstance := db.GetInstance()
	handlers.NewAuthHandler(e, dbInstance)
	handlers.NewEmployeeHandler(e, dbInstance)
	handlers.NewInvitationHandler(e, dbInstance)
	handlers.NewOrganizationHandler(e, dbInstance)
	utils.GetRoles("ankit@gmail.com")
	err := e.Start(":1234")
	if err != nil {
		panic(err)
	}
}
