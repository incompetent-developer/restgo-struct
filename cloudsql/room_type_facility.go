package cloudsql

import "time"

type RoomTypeFacility struct {
	RoomTypeID string    `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType   *RoomType `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	FacilityID string    `json:"facility_id" struct:"facility_id" gorm:"foreignKey:facility_id"`
	Facility   *Facility `json:"facility,omitempty" struct:"facility,omitempty" gorm:"foreignKey:facility_id"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
