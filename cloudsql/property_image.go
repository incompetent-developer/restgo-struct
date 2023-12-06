package cloudsql

import "time"

type PropertyImage struct {
	ID           string    `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID   string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property     *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	ImageUrl     string    `json:"image_url" struct:"image_url"`
	ImageStorage string    `json:"image_storage" struct:"image_storage"`
	PrimaryImage bool      `json:"primary_image" struct:"primary_image"`
	CreatedAt    time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" struct:"updated_at"`
}
