package cloudsql

import "time"

type RoomBed struct {
	RoomID      string     `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room        *Room      `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	TypeOfBedID string     `json:"type_of_bed_id" struct:"type_of_bed_id" gorm:"foreignKey:type_of_bed_id"`
	TypeOfBed   *TypeOfBed `json:"type_of_bed,omitempty" struct:"type_of_bed,omitempty" gorm:"foreignKey:type_of_bed_id"`
	NumberBed   int64      `json:"number_bed" struct:"number_bed" gorm:"default:1"`
	CreatedAt   time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" struct:"updated_at"`
}
