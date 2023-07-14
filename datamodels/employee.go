package datamodels

type Employee struct {
	//ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4(),primaryKey"`
	Email          string `json:"Email" validate:"required" gorm:"primaryKey"`
	Name           string `json:"Name" validate:"required,email"`
	ProfilePicture string `json:"picture"`
	OrgId          int    `json:"OrgId" validate:"required"`
	Role           string `json:"Role"`
	ManagedBy      string
}
type GetEmployee struct {
	Email          string `json:"Email" validate:"required"`
	Name           string `json:"Name" validate:"required,email"`
	ProfilePicture string `json:"picture"`
	ManagedBy      string `json:"ManagedBy,omitempty"`
}
type EditEmployee struct {
	Email     string `json:"Email" validate:"required,email"`
	ManagedBy string `json:"ManagedBy" validate:"required,email"`
}
