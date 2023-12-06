package cloudsql

import (
	"time"
)

//PartnerTransactionBookingDetail structure
type PartnerTransactionBookingDetail struct {
	BookingID      string     `json:"booking_id" structs:"booking_id" gorm:"primary_key;auto_increment:false"`
	RoomID         string     `json:"room_id" structs:"room_id" gorm:"primary_key;auto_increment:false"`
	BasePrice      float64    `json:"base_price" structs:"base_price"`
	OfferPrice     float64    `json:"offer_price" structs:"offer_price"`
	NumberExtraBed int        `json:"number_extra_bed" structs:"number_extra_bed"`
	ExtraBedPrice  float64    `json:"extra_bed_price" structs:"extra_bed_price"`
	DiscountPrice  float64    `json:"discount_price" structs:"discount_price"`
	MarkupPrice    float64    `json:"markup_price" structs:"markup_price"`
	FinalPrice     float64    `json:"final_price" structs:"final_price"`
	CreatedAt      time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID       *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType     *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
