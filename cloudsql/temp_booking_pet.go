package cloudsql

import "time"

type TempBookingPet struct {
	TempBookingID    string          `json:"temp_booking_id" struct:"temp_booking_id" gorm:"foreignKey:temp_booking_id"`
	TempBooking      *TempBooking    `json:"temp_booking,omitempty" struct:"temp_booking,omitempty" gorm:"foreignKey:temp_booking_id"`
	PetTypeOfPriceID *string         `json:"pet_type_of_price_id,omitempty" struct:"pet_type_of_price_id,omitempty"`
	PetTypeOfPrice   *PetTypeOfPrice `json:"pet_type_of_price,omitempty" struct:"pet_type_of_price,omitempty" gorm:"foreignKey:pet_type_of_price_id"`
	NumberOfPet      int64           `json:"number_of_pet" struct:"number_of_pet" `
	MinWeight        *int64          `json:"min_weight,omitempty" struct:"min_weight,omitempty"`
	MaxWeight        *int64          `json:"max_weight,omitempty" struct:"max_weight,omitempty"`
	Price            float64         `json:"price" struct:"price"`
	TotalPrice       float64         `json:"total_price" struct:"total_price"`
	CreatedAt        time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" struct:"updated_at"`
}
