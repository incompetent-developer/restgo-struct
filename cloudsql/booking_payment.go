package cloudsql

import "time"

type BookingPayment struct {
	PaymentHistoryID string          `json:"payment_history_id" struct:"payment_history_id" gorm:"foreignKey:payment_history_id"`
	PaymentHistory   *PaymentHistory `json:"payment_history,omitempty" struct:"payment_history,omitempty" gorm:"foreignKey:payment_history_id"`
	BookingID        *string         `json:"booking_id,omitempty" struct:"booking_id,omitempty" gorm:"foreignKey:booking_id"`
	Booking          *Booking        `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	CreatedAt        time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" struct:"updated_at"`
}
