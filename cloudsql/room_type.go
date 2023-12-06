package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type RoomType struct {
	ID                  string               `json:"id" struct:"id" gorm:"primaryKey"`
	Name                string               `json:"name" struct:"name"`
	PropertyID          string               `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property            *Property            `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	TypeOfRoomID        string               `json:"type_of_room_id" struct:"type_of_room_id" gorm:"foreignKey:type_of_room_id"`
	TypeOfRoom          *TypeOfRoom          `json:"type_of_room,omitempty" struct:"type_of_room,omitempty" gorm:"foreignKey:type_of_room_id"`
	Size                float64              `json:"size" struct:"size"`
	NumberBedroom       int64                `json:"number_bedroom" struct:"number_bedroom"`
	NumberToilet        int64                `json:"number_toilet" struct:"number_toilet"`
	TypeOfViewID        string               `json:"type_of_view_id" struct:"type_of_view_id" gorm:"foreignKey:type_of_view_id"`
	TypeOfView          *TypeOfView          `json:"type_of_view,omitempty" struct:"type_of_view,omitempty" gorm:"foreignKey:type_of_view_id"`
	Balcony             bool                 `json:"balcony" struct:"balcony"`
	Notation            *string              `json:"notation" struct:"notation"`
	NumberAdult         int64                `json:"number_adult" struct:"number_adult"`
	NumberChild         int64                `json:"number_child" struct:"number_child"`
	Babycot             bool                 `json:"babycot" struct:"babycot"`
	NumberBreakfast     int64                `json:"number_breakfast" struct:"number_breakfast"`
	NumberExtraBed      int64                `json:"number_extra_bed" struct:"number_extra_bed"`
	IsAvailabled        bool                 `json:"is_availabled" struct:"is_availabled"`
	IsEnabled           bool                 `json:"is_enabled" struct:"is_enabled"`
	IsSuspended         bool                 `json:"is_suspended" struct:"is_suspended"`
	UpdatedTime         *time.Time           `json:"updated_time" struct:"updated_time"`
	CreatedAt           time.Time            `json:"created_at" struct:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at" struct:"updated_at"`
	DeletedAt           *gorm.DeletedAt      `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	RoomTypeImages      *[]RoomTypeImage     `json:"room_type_images,omitempty" struct:"room_type_images,omitempty" gorm:"foreignKey:room_type_id"`
	RoomTypeBeds        *[]RoomTypeBed       `json:"room_type_beds,omitempty" struct:"room_type_beds,omitempty" gorm:"foreignKey:room_type_id"`
	RoomTypeFacilities  *[]RoomTypeFacility  `json:"room_type_facilities,omitempty" struct:"room_type_facilities,omitempty" gorm:"foreignKey:room_type_id"`
	RoomTypePrice       *RoomTypePrice       `json:"room_type_price,omitempty" struct:"room_type_price,omitempty" gorm:"foreignKey:room_type_id"`
	RoomTypeTemporaries *[]RoomTypeTemporary `json:"room_type_temporaries,omitempty" struct:"room_type_temporaries,omitempty" gorm:"foreignKey:room_type_id"`
	Rooms               *[]Room              `json:"rooms,omitempty" struct:"rooms,omitempty" gorm:"foreignKey:room_type_id"`
}
