package cloudsql

import "time"

type RoomTypeTemporary struct {
	RoomTypeID    string    `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType      *RoomType `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	TemporaryDate time.Time `json:"temporary_date" struct:"temporary_date"`
	CreatedAt     time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" struct:"updated_at"`
}
