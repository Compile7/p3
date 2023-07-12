package utils

import (
	"context"
	"fmt"
	"p3/entities"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

func TokenVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		SocialClientKey := "711686112966-qme1urtstdo0thlv10v0rjv62nelpf5r.apps.googleusercontent.com"
		token := c.Request().Header.Get("Authorization")
		payload, err := idtoken.Validate(context.Background(), token, SocialClientKey)
		if err != nil {

		}
		var SocialClaims entities.SocialTokenClaims
		mapstructure.Decode(payload.Claims, &SocialClaims)

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
		if !SocialClaims.EmailNotVerified() {
			//handle error
			return echo.ErrBadRequest
		}
		c.Set("Email", SocialClaims.Email)
		c.Set("Name", SocialClaims.FullName)
		c.Set("ProfilePicture", SocialClaims.ProfilePicture)
		return next(c)
	}
}
