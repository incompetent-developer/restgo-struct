package cloudsql

import "time"

type TempTransactionBooking struct {
	TempTransactionID  string           `json:"temp_transaction_id" struct:"temp_transaction_id" gorm:"foreignKey:temp_transaction_id"`
	TempTransaction    *TempTransaction `json:"temp_transaction,omitempty" struct:"temp_transaction,omitempty" gorm:"foreignKey:temp_transaction_id"`
	BookingID          string           `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking            *Booking         `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	Channel            *string          `json:"channel,omitempty" struct:"channel,omitempty"`
	Fee                *float64         `json:"fee,omitempty" struct:"fee,omitempty"`
	Vat                *float64         `json:"vat,omitempty" struct:"vat,omitempty"`
	Tax3               *float64         `json:"tax3"`
	CancelPolicyCharge *float64         `json:"cancel_policy_charge,omitempty" struct:"cancel_policy_charge,omitempty"`
	CancelPolicyPrice  *float64         `json:"cancel_policy_price,omitempty" struct:"cancel_policy_price,omitempty"`
	TotalPrice         float64          `json:"total_price" struct:"total_price"`
	CreatedAt          time.Time        `json:"created_at" struct:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at" struct:"updated_at"`
}
