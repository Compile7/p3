package entities

type SocialTokenClaims struct {
	Audience        string `mapstructure:"aud"`
	ClientKey       string `mapstructure:"azp"`
	Email           string `mapstructure:"email"`
	IsEmailVerified bool   `mapstructure:"email_verified"`
	ExpiryTime      int64  `mapstructure:"exp"`
	LastName        string `mapstructure:"family_name"`
	FirstName       string `mapstructure:"given_name"`
	IssueTime       int64  `mapstructure:"iat"`
	Issuer          string `mapstructure:"iss"`
	Identifier      string `mapstructure:"jti"`
	FullName        string `mapstructure:"name"`
	NotBeforeTime   int64  `mapstructure:"nbf"`
	ProfilePicture  string `mapstructure:"picture"`
}

func (c *SocialTokenClaims) IsNotExpired(now int64) bool {
	if c.ExpiryTime == 0 {
		return false
	}

	return now <= c.ExpiryTime
}

func (c *SocialTokenClaims) IsIssueTimeValid(now int64) bool {
	if c.IssueTime == 0 {
		return false
	}

	return now >= c.IssueTime
}

func (c *SocialTokenClaims) IsAudienceValid(GoogleClientKey string) bool {
	if len(c.Audience) == 0 {
		return false
	}

	return c.Audience == GoogleClientKey
}

func (c *SocialTokenClaims) IsIssuerValid() bool {
	if len(c.Issuer) == 0 {
		return false
	}

	return c.Issuer == "https://accounts.google.com"
}

func (c *SocialTokenClaims) DoesApiKeyMatch(GoogleClientKey string) bool {
	return GoogleClientKey == c.ClientKey
}

func (c *SocialTokenClaims) EmailNotVerified() bool {

	return !c.IsEmailVerified
}
