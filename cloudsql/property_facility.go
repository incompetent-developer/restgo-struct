package cloudsql

import "time"

type PropertyFacility struct {
	PropertyID string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property   *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	FacilityID string    `json:"facility_id" struct:"facility_id" gorm:"foreignKey:facility_id"`
	Facility   *Facility `json:"facility,omitempty" struct:"facility,omitempty" gorm:"foreignKey:facility_id"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
