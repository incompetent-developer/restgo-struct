package cloudsql

import "time"

type BookingPaymentType struct {
	ID        string    `json:"id" struct:"id" gorm:"primaryKey"`
	Name      string    `json:"name" struct:"name"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
