package cloudsql

import "time"

type TempBookingPayment struct {
	TempPaymentHistoryID string              `json:"temp_payment_history_id" struct:"temp_payment_history_id" gorm:"foreignKey:temp_payment_history_id"`
	TempPaymentHistory   *TempPaymentHistory `json:"temp_payment_history,omitempty" struct:"temp_payment_history,omitempty" gorm:"foreignKey:temp_payment_history_id"`
	TempBookingID        *string             `json:"temp_booking_id,omitempty" struct:"temp_booking_id,omitempty" gorm:"foreignKey:temp_booking_id"`
	TempBooking          *TempBooking        `json:"temp_booking,omitempty" struct:"temp_booking,omitempty" gorm:"foreignKey:temp_booking_id"`
	// BookingID            *string             `json:"booking_id,omitempty" struct:"booking_id,omitempty" gorm:"foreignKey:booking_id"`
	// Booking              *Booking            `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
