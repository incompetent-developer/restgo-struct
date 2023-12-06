package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type TypeOfView struct {
	ID          string          `json:"id" struct:"id" gorm:"primaryKey"`
	Name        string          `json:"name" struct:"name"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
