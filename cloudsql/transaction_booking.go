package cloudsql

import "time"

type TransactionBooking struct {
	TransactionID      string       `json:"transaction_id" struct:"transaction_id" gorm:"foreignKey:transaction_id"`
	Transaction        *Transaction `json:"transaction,omitempty" struct:"transaction,omitempty" gorm:"foreignKey:transaction_id"`
	BookingID          string       `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking            *Booking     `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	Channel            *string      `json:"channel" struct:"channel"`
	Fee                *float64     `json:"fee" struct:"fee"`
	Vat                *float64     `json:"vat" struct:"vat"`
	Tax3               *float64     `json:"tax3"`
	CancelPolicyCharge *float64     `json:"cancel_policy_charge" struct:"cancel_policy_charge"`
	CancelPolicyPrice  *float64     `json:"cancel_policy_price" struct:"cancel_policy_price"`
	TotalPrice         float64      `json:"total_price" struct:"total_price"`
	CreatedAt          time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at" struct:"updated_at"`
}
