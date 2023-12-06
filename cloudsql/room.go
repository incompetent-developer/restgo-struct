package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID              string          `json:"id" struct:"id" gorm:"primaryKey"`
	Name            string          `json:"name" struct:"name"`
	PropertyID      string          `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property        *Property       `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	RoomTypeID      string          `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType        *RoomType       `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	Floor           *string         `json:"floor,omitempty" struct:"floor,omitempty"`
	Size            float64         `json:"size" struct:"size"`
	NumberBedroom   int64           `json:"number_bedroom" struct:"number_bedroom" gorm:"default:0"`
	NumberToilet    int64           `json:"number_toilet" struct:"number_toilet" gorm:"default:1"`
	TypeOfViewID    string          `json:"type_of_view_id" struct:"type_of_view_id" gorm:"foreignKey:type_of_view_id"`
	TypeOfView      *TypeOfView     `json:"type_of_view,omitempty" struct:"type_of_view,omitempty" gorm:"foreignKey:type_of_view_id"`
	Balcony         bool            `json:"balcony" struct:"balcony"`
	Notation        *string         `json:"notation,omitempty" struct:"notation,omitempty"`
	NumberAdult     int64           `json:"number_adult" struct:"number_adult" gorm:"default:1"`
	NumberChild     int64           `json:"number_child" struct:"number_child" gorm:"default:0"`
	Babycot         bool            `json:"babycot" struct:"babycot"`
	NumberBreakfast int64           `json:"number_breakfast" struct:"number_breakfast" gorm:"default:0"`
	NumberExtraBed  int64           `json:"number_extra_bed" struct:"number_extra_bed" gorm:"default:0"`
	IsLayout        bool            `json:"is_layout" struct:"is_layout"`
	IsAvailabled    bool            `json:"is_availabled" struct:"is_availabled"`
	IsEnabled       bool            `json:"is_enabled" struct:"is_enabled"`
	IsSuspended     bool            `json:"is_suspended" struct:"is_suspended"`
	UpdatedTime     *time.Time      `json:"updated_time,omitempty" struct:"updated_time,omitempty"`
	CreatedAt       time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	RoomImages      *[]RoomImage    `json:"room_images,omitempty" struct:"room_images,omitempty" gorm:"foreignKey:room_id"`
	RoomBeds        *[]RoomBed      `json:"room_beds,omitempty" struct:"room_beds,omitempty" gorm:"foreignKey:room_id"`
	RoomFacilities  *[]RoomFacility `json:"room_facilities,omitempty" struct:"room_facilities,omitempty" gorm:"foreignKey:room_id"`
	RoomPrice       *RoomPrice      `json:"room_price,omitempty" struct:"room_price,omitempty" gorm:"foreignKey:room_id"`
	RoomLayout      *RoomLayout     `json:"room_layout,omitempty" struct:"room_layout,omitempty" gorm:"foreignKey:room_id"`
}
