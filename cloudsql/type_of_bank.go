package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type TypeOfBank struct {
	ID          string          `json:"id" struct:"id" gorm:"primaryKey"`
	NameTH      string          `json:"name_th" struct:"name_th"`
	NameEN      string          `json:"name_en" struct:"name_en"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
