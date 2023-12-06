package cloudsql

import "time"

type RoomLayout struct {
	RoomID    string    `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room      *Room     `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	XAxis     float64   `json:"x_axis" struct:"x_axis" gorm:"default:0"`
	YAxis     float64   `json:"y_axis" struct:"y_axis" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
