package cloudsql

import "time"

type RoomImage struct {
	ID           string    `json:"id" struct:"id"`
	RoomID       string    `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room         *Room     `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	ImageUrl     string    `json:"image_url" struct:"image_url"`
	ImageStorage string    `json:"image_storage" struct:"image_storage"`
	PrimaryImage bool      `json:"primary_image" struct:"primary_image"`
	CreatedAt    time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" struct:"updated_at"`
}
