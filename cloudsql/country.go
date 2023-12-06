package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID          int             `json:"id" struct:"id" gorm:"primaryKey"`
	NameTH      string          `json:"name_th" struct:"name_th"`
	NameEN      string          `json:"name_en" struct:"name_en"`
	ContinentID int             `json:"continent_id" struct:"continent_id" gorm:"foreignKey:continent_id"`
	Continent   *Continent      `json:"continent,omitempty" struct:"continent,omitempty" gorm:"foreignKey:continent_id"`
	IsSuspended bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
