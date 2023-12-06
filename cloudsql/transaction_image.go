package cloudsql

import "time"

type TransactionImage struct {
	ID               string       `json:"id" struct:"id" gorm:"primaryKey"`
	TransactionID    string       `json:"transaction_id" struct:"transaction_id" gorm:"foreignKey:transaction_id"`
	Transaction      *Transaction `json:"transaction,omitempty" struct:"transaction,omitempty" gorm:"foreignKey:transaction_id"`
	ImageUrl         string       `json:"image_url" struct:"image_url"`
	ImageStorage     string       `json:"image_storage" struct:"image_storage"`
	DateTimeTransfer time.Time    `json:"date_time_transfer" struct:"date_time_transfer"`
	Note             *string      `json:"note" struct:"note"`
	CreatedAt        time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" struct:"updated_at"`
}
