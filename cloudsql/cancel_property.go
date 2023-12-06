package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type CancelPolicy struct {
	ID          string          `json:"id" struct:"id" gorm:"primaryKey"`
	NameTH      string          `json:"name_th" struct:"name_th"`
	NameEN      string          `json:"name_en" struct:"name_en"`
	AmountDay   int64           `json:"amount_day" struct:"amount_day" gorm:"default:0"`
	Charge      float64         `json:"charge" struct:"charge" gorm:"default:0"`
	NoShow      float64         `json:"no_show" struct:"no_show" gorm:"default:0"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
