package cloudsql

import "time"

type PropertyPet struct {
	ID         string    `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property   *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	MinWeight  *int64    `json:"min_weight,omitempty" struct:"min_weight,omitempty"`
	MaxWeight  *int64    `json:"max_weight,omitempty" struct:"max_weight,omitempty"`
	MaximumPet int64     `json:"maximum_pet" struct:"maximum_pet"`
	Price      float64   `json:"price" struct:"price"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
