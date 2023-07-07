package errors

type CustomError struct {
	status int
	msg    string
}

var (
	ErrOrgnaisationRequired = CustomError{status: 403, msg: "You are not a part of any organisation"}
	ErrInvalidToken         = CustomError{status: 403, msg: "invalid token"}
)
