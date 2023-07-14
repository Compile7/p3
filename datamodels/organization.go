package datamodels

import (
	"time"
)

type Organization struct {
	ID        int    `json:"id,omitempty" gorm:"primaryKey"`
	Name      string `json:"Name" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string `json:"CreatedBy,omitempty"`
	IsActive  bool   `json:"IsActive,omitempty"`
	IsDeleted bool   `json:"IsDeleted,omitempty"`
}
type EditOrganization struct {
	Name string `json:"Name" validate:"required"`
}
type GetOrganization struct {
	//ID        int    `gorm:"primaryKey"`
	Name      string `json:"Name" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//CreatedBy string `json:"CreatedBy"`
	//IsActive  bool   `json:"IsActive"`
	//IsDeleted bool   `json:"IsDeleted"`
}
