package cloudsql

import "time"

type RoomTypePrice struct {
	ID                    string                  `json:"id" struct:"id" gorm:"primaryKey"`
	RoomTypeID            string                  `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType              *RoomType               `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	StartDate             *time.Time              `json:"start_date" struct:"start_date"`
	EndDate               *time.Time              `json:"end_date" struct:"end_date"`
	Forever               bool                    `json:"forever" struct:"forever"`
	BasePrice             *float64                `json:"base_price" struct:"base_price"`
	ExtraBedPrice         *float64                `json:"extra_bed_price" struct:"extra_bed_price"`
	TypeOfFakePriceID     *string                 `json:"type_of_fake_price_id" struct:"type_of_fake_price_id" gorm:"foreignKey:type_of_fake_price_id"`
	FakePrice             *float64                `json:"fake_price" struct:"fake_price"`
	IsEnabled             bool                    `json:"is_enabled" struct:"is_enabled"`
	RoomTypePriceSettings *[]RoomTypePriceSetting `json:"room_type_price_settings,omitempty" struct:"room_type_price_settings,omitempty" gorm:"foreignKey:room_type_price_id"`
	RoomTypePriceOffers   *[]RoomTypePriceOffer   `json:"room_type_price_offers,omitempty" struct:"room_type_price_offers,omitempty" gorm:"foreignKey:room_type_price_id"`
	RoomTypePriceHolidays *[]RoomTypePriceHoliday `json:"room_type_price_holidays,omitempty" struct:"room_type_price_holidays,omitempty" gorm:"foreignKey:room_type_price_id"`
	CreatedAt             time.Time               `json:"created_at" struct:"created_at"`
	UpdatedAt             time.Time               `json:"updated_at" struct:"updated_at"`
}
