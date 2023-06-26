package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"p3/db"
	handlers "p3/handler"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	dbInstance := db.GetInstance()
	handlers.NewEmployeeHandler(e, dbInstance)
	handlers.NewOrganizationHandler(e, dbInstance)
	err := e.Start(":1234")
	if err != nil {
		panic(err)
	}
}
