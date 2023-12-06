package cloudsql

import "time"

type FineBooking struct {
	FineID          string    `json:"fine_id" struct:"fine_id" gorm:"foreignKey:fine_id"`
	Fine            *Fine     `json:"fine,omitempty" struct:"fine,omitempty" gorm:"foreignKey:fine_id"`
	BookingID       string    `json:"booking_id,omitempty" struct:"booking_id,omitempty" gorm:"foreignKey:booking_id"`
	Booking         *Booking  `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	TotalPrice      float64   `json:"total_price" struct:"total_price"`
	FineCharge      float64   `json:"fine_charge" struct:"fine_charge"`
	FineChargePrice float64   `json:"fine_charge_price" struct:"fine_charge_price"`
	TotalFinePrice  float64   `json:"total_fine_price" struct:"total_fine_price"`
	CreatedAt       time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" struct:"updated_at"`
}
