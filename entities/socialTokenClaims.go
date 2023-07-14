package entities

type SocialTokenClaims struct {
	Audience        string `json:"aud"`
	ClientKey       string `json:"azp"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"email_verified"`
	ExpiryTime      int64  `json:"exp"`
	LastName        string `json:"family_name"`
	FirstName       string `json:"given_name"`
	IssueTime       int64  `json:"iat"`
	Issuer          string `json:"iss"`
	Identifier      string `json:"jti"`
	FullName        string `json:"name"`
	NotBeforeTime   int64  `json:"nbf"`
	ProfilePicture  string `json:"picture"`
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
