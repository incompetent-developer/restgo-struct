package cloudsql

import "time"

type RoomTypeBed struct {
	RoomTypeID  string     `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType    *RoomType  `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	TypeOfBedID string     `json:"type_of_bed_id" struct:"type_of_bed_id" gorm:"foreignKey:type_of_bed_id"`
	TypeOfBed   *TypeOfBed `json:"type_of_bed,omitempty" struct:"type_of_bed,omitempty" gorm:"foreignKey:type_of_bed_id"`
	NumberBed   int64      `json:"number_bed" struct:"number_bed" gorm:"default:1"`
	CreatedAt   time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" struct:"updated_at"`
}
