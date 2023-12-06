package cloudsql

import "time"

type TempBookingBasket struct {
	BasketID      string       `json:"basket_id" struct:"basket_id" gorm:"foreignKey:basket_id"`
	Basket        *Basket      `json:"basket,omitempty" struct:"basket,omitempty" gorm:"foreignKey:basket_id"`
	TempBookingID string       `json:"temp_booking_id" struct:"temp_booking_id" gorm:"foreignKey:temp_booking_id"`
	TempBooking   *TempBooking `json:"temp_booking,omitempty" struct:"temp_booking,omitempty" gorm:"foreignKey:temp_booking_id"`
	CreatedAt     time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" struct:"updated_at"`
}
