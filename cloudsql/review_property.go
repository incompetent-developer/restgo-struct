package cloudsql

import "time"

type ReviewProperty struct {
	ReviewID   string    `json:"review_id" struct:"review_id" gorm:"primaryKey"`
	Review     *Review   `json:"review" struct:"review" gorm:"primaryKey:review_id"`
	PropertyID string    `json:"property_id" struct:"property_id" gorm:"primaryKey"`
	Property   *Review   `json:"property" struct:"property" gorm:"primaryKey:property_id"`
	Message    string    `json:"message" struct:"message"`
	IsEnabled  bool      `json:"is_enabled" struct:"is_enabled"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
