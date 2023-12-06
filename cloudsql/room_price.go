package cloudsql

import "time"

type RoomPrice struct {
	ID                string              `json:"id" struct:"id" gorm:"primaryKey"`
	RoomID            string              `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room              *Room               `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	StartDate         *time.Time          `json:"start_date,omitempty" struct:"start_date,omitempty"`
	EndDate           *time.Time          `json:"end_date,omitempty" struct:"end_date,omitempty"`
	Forever           bool                `json:"forever" struct:"forever"`
	BasePrice         *float64            `json:"base_price,omitempty" struct:"base_price,omitempty"`
	ExtraBedPrice     *float64            `json:"extra_bed_price,omitempty" struct:"extra_bed_price,omitempty"`
	TypeOfFakePriceID *string             `json:"type_of_fake_price_id" struct:"type_of_fake_price_id" gorm:"foreignKey:type_of_fake_price_id"`
	FakePrice         *float64            `json:"fake_price,omitempty" struct:"fake_price,omitempty"`
	IsEnabled         bool                `json:"is_enabled" struct:"is_enabled"`
	RoomPriceSettings *[]RoomPriceSetting `json:"room_price_settings,omitempty" struct:"room_price_settings,omitempty" gorm:"foreignKey:room_price_id"`
	RoomPriceHolidays *[]RoomPriceHoliday `json:"room_price_holidays,omitempty" struct:"room_price_holidays,omitempty" gorm:"foreignKey:room_price_id"`
	CreatedAt         time.Time           `json:"created_at" struct:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at" struct:"updated_at"`
}
