package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Landmark struct {
	ID               string          `json:"id" struct:"id" gorm:"primaryKey"`
	Name             string          `json:"name" struct:"name"`
	Latitude         string          `json:"latitude" struct:"latitude"`
	Longtitude       string          `json:"longtitude" struct:"longtitude"`
	TypeOfLandmarkID string          `json:"type_of_landmark_id" struct:"type_of_landmark_id" gorm:"foreignKey:type_of_landmark_id"`
	TypeOfLandmark   *TypeOfLandmark `json:"type_of_landmark,omitempty" struct:"type_of_landmark,omitempty" gorm:"foreignKey:type_of_landmark_id"`
	IsEnabled        bool            `json:"is_enabled" struct:"is_enabled"`
	IsSuspended      bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt        time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt        *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
