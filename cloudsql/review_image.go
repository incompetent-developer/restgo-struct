package cloudsql

import "time"

type ReviewImage struct {
	ID           string    `json:"id" struct:"id" gorm:"primaryKey"`
	ReviewID     string    `json:"review_id" struct:"review_id" gorm:"foreignKey:review_id"`
	Review       *Review   `json:"review,omitempty" struct:"review,omitempty" gorm:"foreignKey:review_id"`
	ImageUrl     string    `json:"image_url" struct:"image_url"`
	ImageStorage string    `json:"image_storage" struct:"image_storage"`
	CreatedAt    time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" struct:"updated_at"`
}
