package cloudsql

import "time"

type RoomPriceHoliday struct {
	RoomPriceID string     `json:"room_price_id" struct:"room_price_id" gorm:"foreignKey:room_price_id"`
	RoomPrice   *RoomPrice `json:"room_price,omitempty" struct:"room_price,omitempty" gorm:"foreignKey:room_price_id"`
	Holiday     time.Time  `json:"holiday" struct:"holiday"`
	Price       float64    `json:"price" struct:"price"`
	CreatedAt   time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" struct:"updated_at"`
}
