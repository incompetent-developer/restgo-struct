package cloudsql

import "time"

type RoomTemporary struct {
	RoomID        string    `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room          *Room     `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	TemporaryDate time.Time `json:"temporary_date" struct:"temporary_date"`
	CreatedAt     time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" struct:"updated_at"`
}
