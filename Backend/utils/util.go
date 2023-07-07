package utils

import (
	"context"
	"fmt"
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
