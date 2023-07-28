package cErr

type CustomError struct {
	Status  int    `json:"Status,omitempty"`
	Message string `json:"Message,omitempty"`
}

var (
	OrgnaisationRequired    = CustomError{Status: 403, Message: "You are not a part of any organisation"}
	NotAllowed              = CustomError{Status: 403, Message: "You don't have required permissions to perform this operation"}
	InvalidToken            = CustomError{Status: 403, Message: "invalid token"}
	InvalidEmail            = CustomError{Status: 403, Message: "invalid email"}
	TokenMissing            = CustomError{Status: 403, Message: "Missing auth token"}
	Mapping                 = CustomError{Status: 403, Message: "Mapping error"}
	DataNotFound            = CustomError{Status: 403, Message: "Data not found"}
	OrgAlreadyExist         = CustomError{Status: 403, Message: "You have already created the organization"}
	InvalidPostBody         = CustomError{Status: 400, Message: "Invalid post body"}
	ConnectionError         = CustomError{Status: 400, Message: "Connection Error"}
	AccessDenied            = CustomError{Status: 403, Message: "Update access denied"}
	EmployeeNotFound        = CustomError{Status: 403, Message: "Employee not found"}
	ManagerNotFound         = CustomError{Status: 403, Message: "Manager not found"}
	EmailAndManageByNotSame = CustomError{Status: 403, Message: "Email and ManageBy should not be same"}
	UpdateError             = CustomError{Status: 403, Message: "Update error"}
)
