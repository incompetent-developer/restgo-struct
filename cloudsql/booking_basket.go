package cloudsql

import "time"

type BookingBasket struct {
	BasketID  string    `json:"basket_id" struct:"basket_id" gorm:"foreignKey:basket_id"`
	Basket    *Basket   `json:"basket,omitempty" struct:"basket,omitempty" gorm:"foreignKey:basket_id"`
	BookingID string    `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking   *Booking  `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
