package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	ID          int             `json:"id" struct:"id" gorm:"primaryKey"`
	NameTH      string          `json:"name_th" struct:"name_th"`
	NameEN      string          `json:"name_en" struct:"name_en"`
	CountryID   int             `json:"country_id" struct:"country_id" gorm:"foreignKey:country_id"`
	Country     *Country        `json:"country,omitempty" struct:"country,omitempty" gorm:"foreignKey:country_id"`
	RegionID    int             `json:"region_id" struct:"region_id" gorm:"foreignKey:region_id"`
	Region      *Region         `json:"region,omitempty" struct:"region,omitempty" gorm:"foreignKey:region_id"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
