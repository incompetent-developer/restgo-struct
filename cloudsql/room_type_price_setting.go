package cloudsql

import "time"

type RoomTypePriceSetting struct {
	RoomTypePriceID string         `json:"room_type_price_id" struct:"room_type_price_id" gorm:"foreignKey:room_type_price_id"`
	RoomTypePrice   *RoomTypePrice `json:"room_type_price,omitempty" struct:"room_type_price,omitempty" gorm:"foreignKey:room_type_price_id"`
	DayOfWeekID     string         `json:"day_of_week_id" struct:"day_of_week_id" gorm:"foreignKey:day_of_week_id"`
	DayOfWeek       *DayOfWeek     `json:"price_setting,omitempty" struct:"price_setting,omitempty" gorm:"foreignKey:day_of_week_id"`
	BasePrice       float64        `json:"base_price" struct:"base_price" gorm:"default:0"`
	IsEnabled       bool           `json:"is_enabled" struct:"is_enabled"`
	CreatedAt       time.Time      `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" struct:"updated_at"`
}
