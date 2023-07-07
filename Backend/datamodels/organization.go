package datamodels

import (
	"time"
)

type Organization struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `json:"Name" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string `json:"CreatedBy"`
	IsActive  bool   `json:"IsActive"`
	IsDeleted bool   `json:"IsDeleted"`
}
