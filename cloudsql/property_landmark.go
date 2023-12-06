package cloudsql

import "time"

type PropertyLandmark struct {
	PropertyID string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property   *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	LandmarkID string    `json:"landmark_id" struct:"landmark_id" gorm:"foreignKey:landmark_id"`
	Landmark   *Landmark `json:"landmark,omitempty" struct:"landmark,omitempty" gorm:"foreignKey:landmark_id"`
	Distance   float64   `json:"float" struct:"float" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
