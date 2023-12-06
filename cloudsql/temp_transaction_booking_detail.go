package cloudsql

import "time"

type TempTransactionBookingDetail struct {
	TempTransactionID string           `json:"temp_transaction_id" struct:"temp_transaction_id" gorm:"primaryKey"`
	TempTransaction   *TempTransaction `json:"temp_transaction,omitempty" struct:"temp_transaction,omitempty" gorm:"foreignKey:temp_transaction_id"`
	BookingID         string           `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking           *Booking         `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	RoomTypeID        string           `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType          *Booking         `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	AmountRoom        int64            `json:"amount_room" struct:"amount_room"`
	AmountNight       int64            `json:"amount_night" struct:"amount_night"`
	BasePrice         float64          `json:"base_price" struct:"base_price"`
	AddonPrice        *float64         `json:"addon_price" struct:"addon_price"`
	PricePerOffer     *float64         `json:"price_per_offer,omitempty" struct:"price_per_offer,omitempty"`
	Offer             *float64         `json:"offer,omitempty" struct:"offer,omitempty"`
	PetPrice          *float64         `json:"pet_price,omitempty" struct:"pet_price,omitempty"`
	PricePerExtraBed  *float64         `json:"price_per_extra_bed,omitempty" struct:"price_per_extra_bed,omitempty"`
	NumberExtraBed    *int64           `json:"number_extra_bed,omitempty" struct:"number_extra_bed,omitempty"`
	ExtraBedPrice     *float64         `json:"extra_bed_price,omitempty" struct:"extra_bed_price,omitempty"`
	Discount          *float64         `json:"discount,omitempty" struct:"discount,omitempty"`
	TotalPrice        float64          `json:"total_price" struct:"total_price"`
	CreatedAt         time.Time        `json:"created_at" struct:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at" struct:"updated_at"`
}
