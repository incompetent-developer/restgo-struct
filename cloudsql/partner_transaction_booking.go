package cloudsql

import (
	"time"
)

//PartnerTransactionBooking structure
type PartnerTransactionBooking struct {
	BookingID          string     `json:"booking_id" structs:"booking_id" gorm:"primary_key;auto_increment:false"`
	Booking            *Booking   `json:"booking,omitempty" structs:"booking,omitempty"`
	TotalBasePrice     float64    `json:"total_base_price" structs:"total_base_price"`
	TotalOfferPrice    float64    `json:"total_offer_price" structs:"total_offer_price"`
	TotalExtraBedPrice float64    `json:"total_extra_bed_price" structs:"total_extra_bed_price"`
	TotalDiscountPrice float64    `json:"total_discount_price" structs:"total_discount_price"`
	TotalMarkupPrice   float64    `json:"total_markup_price" structs:"total_markup_price"`
	PetPrice           float64    `json:"pet_price" structs:"pet_price"`
	Vat                float64    `json:"vat" structs:"vat"`
	NetPrice           float64    `json:"net_price" structs:"net_price"`
	TotalPrice         float64    `json:"total_price" structs:"total_price"`
	VatRate            float64    `json:"vat_rate" structs:"vat_rate"`
	FeeRate            float64    `json:"fee_rate" structs:"fee_rate"`
	WithholdingRate    float64    `json:"withholding_rate" structs:"withholding_rate"`
	MarkupRate         float64    `json:"markup_rate" structs:"markup_rate"`
	FinalPrice         float64    `json:"final_price" structs:"final_price"`
	CreatedAt          time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID           *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType         *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
