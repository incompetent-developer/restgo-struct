package cloudsql

import "time"

type TransactionBookingDetail struct {
	TransactionID    string       `json:"transaction_id" struct:"transaction_id" gorm:"primaryKey"`
	Transaction      *Transaction `json:"transaction,omitempty" struct:"transaction,omitempty" gorm:"foreignKey:transaction_id"`
	BookingID        string       `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking          *Booking     `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	RoomTypeID       string       `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType         *Booking     `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	AmountRoom       int64        `json:"amount_room" struct:"amount_room"`
	AmountNight      int64        `json:"amount_night" struct:"amount_night"`
	BasePrice        float64      `json:"base_price" struct:"base_price"`
	AddonPrice       *float64     `json:"addon_price" struct:"addon_price"`
	PricePerOffer    *float64     `json:"price_per_offer" struct:"price_per_offer"`
	Offer            *float64     `json:"offer" struct:"offer"`
	PetPrice         *float64     `json:"pet_price" struct:"pet_price"`
	PricePerExtraBed *float64     `json:"price_per_extra_bed" struct:"price_per_extra_bed"`
	NumberExtraBed   *int64       `json:"number_extra_bed" struct:"number_extra_bed"`
	ExtraBedPrice    *float64     `json:"extra_bed_price" struct:"extra_bed_price"`
	Discount         *float64     `json:"discount" struct:"discount"`
	TotalPrice       float64      `json:"total_price" struct:"total_price"`
	CreatedAt        time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" struct:"updated_at"`
}
