package cloudsql

import "time"

type RoomTypePriceHoliday struct {
	RoomTypePriceID string         `json:"room_type_price_id" struct:"room_type_price_id" gorm:"foreignKey:room_type_price_id"`
	RoomTypePrice   *RoomTypePrice `json:"room_type_price,omitempty" struct:"room_type_price,omitempty" gorm:"foreignKey:room_type_price_id"`
	Holiday         time.Time      `json:"holiday" struct:"holiday"`
	Price           float64        `json:"price" struct:"price"`
	CreatedAt       time.Time      `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" struct:"updated_at"`
}
