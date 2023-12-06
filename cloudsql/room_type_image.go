package cloudsql

import "time"

type RoomTypeImage struct {
	ID           string    `json:"id" struct:"id"`
	RoomTypeID   string    `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType     *RoomType `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	ImageUrl     string    `json:"image_url" struct:"image_url"`
	ImageStorage string    `json:"image_storage" struct:"image_storage"`
	PrimaryImage bool      `json:"primary_image" struct:"primary_image"`
	CreatedAt    time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" struct:"updated_at"`
}
