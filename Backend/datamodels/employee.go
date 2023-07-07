package datamodels

import "github.com/google/uuid"

type Employee struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email string    `json:"Email" validate:"required"`
	Name  string    `json:"Name" validate:"required,email"`
	OrgId int       `json:"OrgId" validate:"required"`
}
