package err

type CustomError struct {
	Status  int    `json:"Status,omitempty"`
	Message string `json:"Message,omitempty"`
}

var (
	ErrOrgnaisationRequired = CustomError{Status: 403, Message: "You are not a part of any organisation"}
	ErrInvalidToken         = CustomError{Status: 403, Message: "invalid token"}
)
