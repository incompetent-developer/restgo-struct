package cloudsql

import "time"

type RoomFacility struct {
	RoomID     string    `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room       *Room     `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	FacilityID string    `json:"facility_id" struct:"facility_id" gorm:"foreignKey:facility_id"`
	Facility   *Facility `json:"facility,omitempty" struct:"facility,omitempty" gorm:"foreignKey:facility_id"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
