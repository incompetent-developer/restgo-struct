package cloudsql

import "time"

type FineImage struct {
	ID               string    `json:"id" struct:"id"`
	FineID           string    `json:"fine_id" struct:"fine_id" gorm:"foreignKey:fine_id"`
	Fine             *Fine     `json:"fine,omitempty" struct:"fine,omitempty" gorm:"foreignKey:fine_id"`
	DateTimeTransfer time.Time `json:"date_time_transfer" struct:"date_time_transfer"`
	TransferPrice    float64   `json:"transafer_price" struct:"transafer_price"`
	ImageUrl         string    `json:"image_url" struct:"image_url"`
	ImageStorage     string    `json:"image_storage" struct:"image_storage"`
	CreatedAt        time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" struct:"updated_at"`
}
