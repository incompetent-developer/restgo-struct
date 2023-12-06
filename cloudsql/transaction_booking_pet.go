package cloudsql

import "time"

type TransactionBookingPet struct {
	TransactionID    string          `json:"transaction_id" struct:"transaction_id" gorm:"primaryKey"`
	Transaction      *Transaction    `json:"transaction,omitempty" struct:"transaction,omitempty" gorm:"foreignKey:transaction_id"`
	BookingID        string          `json:"booking_id" struct:"booking_id"`
	Booking          *Booking        `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	PetTypeOfPriceID *string         `json:"pet_type_of_price_id" struct:"pet_type_of_price_id"`
	PetTypeOfPrice   *PetTypeOfPrice `json:"pet_type_of_price,omitempty" struct:"pet_type_of_price,omitempty" gorm:"foreignKey:pet_type_of_price_id"`
	NumberOfPet      int64           `json:"number_of_pet" struct:"number_of_pet" `
	MinWeight        *int64          `json:"min_weight" struct:"min_weight"`
	MaxWeight        *int64          `json:"max_weight" struct:"max_weight"`
	PricePerPet      float64         `json:"price_per_pet" struct:"price_per_pet"`
	TotalPrice       float64         `json:"total_price" struct:"total_price"`
	CreatedAt        time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" struct:"updated_at"`
}
