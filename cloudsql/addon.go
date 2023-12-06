package cloudsql

import "time"

type Addon struct {
	ID              string    `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID      *string   `json:"property_id,omitempty" struct:"property_id,omitempty" gorm:"foreignKey:property_id,omitempty"`
	Property        *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	RoomTypeID      *string   `json:"room_type_id,omitempty" struct:"room_type_id,omitempty" gorm:"foreignKey:room_type_id,omitempty"`
	RoomType        *RoomType `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	RoomID          *string   `json:"room_id,omitempty" struct:"room_id,omitempty" gorm:"foreignKey:room_id,omitempty"`
	Room            *Room     `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	Compensation    float64   `json:"compensation" struct:"compensation" gorm:"default:0"`
	CompensationMin float64   `json:"compensation_min" struct:"compensation_min" gorm:"default:0"`
	CompensationMax float64   `json:"compensation_max" struct:"compensation_max" gorm:"default:0"`
	IsDefault       bool      `json:"is_default" struct:"is_default"`
	CreatedAt       time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" struct:"updated_at"`
}
