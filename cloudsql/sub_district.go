package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type SubDistrict struct {
	ID          int             `json:"id" struct:"id" gorm:"primaryKey"`
	NameTH      string          `json:"name_th" struct:"name_th"`
	NameEN      string          `json:"name_en" struct:"name_en"`
	DistrictID  int             `json:"district_id" struct:"district_id" gorm:"foreignKey:district_id"`
	District    *District       `json:"district,omitempty" struct:"district,omitempty" gorm:"foreignKey:district_id"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
