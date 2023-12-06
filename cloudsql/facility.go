package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	ID                   string              `json:"id" struct:"id" gorm:"primaryKey"`
	Name                 string              `json:"name" struct:"name"`
	CategoryOfFacilityID string              `json:"category_of_facility_id" struct:"category_of_facility_id" gorm:"foreignKey:category_of_facility_id"`
	CategoryOfFacility   *CategoryOfFacility `json:"category_of_facility,omitempty" struct:"category_of_facility,omitempty" gorm:"foreignKey:category_of_facility_id"`
	TypeOfFacilityID     string              `json:"type_of_facility_id" struct:"type_of_facility_id" gorm:"foreignKey:type_of_facility_id"`
	TypeOfFacility       *TypeOfFacility     `json:"type_of_facility,omitempty" struct:"type_of_facility,omitempty" gorm:"foreignKey:type_of_facility_id"`
	IsEnabled            bool                `json:"is_enabled" struct:"is_enabled"`
	IsSuspended          bool                `json:"is_suspended" struct:"is_suspended"`
	CreatedAt            time.Time           `json:"created_at" struct:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at" struct:"updated_at"`
	DeletedAt            *gorm.DeletedAt     `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
