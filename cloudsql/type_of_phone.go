package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type TypeOfPhone struct {
	ID          string          `json:"id" struct:"id" gorm:"primaryKey"`
	Country     string          `json:"country" struct:"country"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
