package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type AccountSale struct {
	ID             string          `json:"id" struct:"id" gorm:"primaryKey"`
	FirstnameTH    string          `json:"firstname_th" struct:"firstname_th"`
	LastnameTH     string          `json:"lastname_th" struct:"lastname_th"`
	FirstnameEN    *string         `json:"firstname_en" struct:"firstname_en"`
	LastnameEN     *string         `json:"lastname_en" struct:"lastname_en"`
	Position       *string         `json:"position" struct:"position"`
	TypeOfPhoneID  string          `json:"type_of_phone_id" struct:"type_of_phone_id" gorm:"foreignKey:type_of_phone_id"`
	TypeOfPhone    *TypeOfPhone    `json:"type_of_phone,omitempty" struct:"type_of_phone,omitempty" gorm:"foreignKey:type_of_phone_id"`
	Phone          string          `json:"phone" struct:"phone"`
	Email          string          `json:"email" struct:"email"`
	ImageUrl       string          `json:"image_url" struct:"image_url"`
	ImageStorage   string          `json:"image_storage" struct:"image_storage"`
	QrImageUrl     string          `json:"qr_image_url" struct:"qr_image_url"`
	QrImageStorage string          `json:"qr_image_storage" struct:"qr_image_storage"`
	Region         *string         `json:"region" struct:"region"`
	Province       *string         `json:"province" struct:"province"`
	IsEnabled      bool            `json:"is_enabled" struct:"is_enabled"`
	IsSuspended    bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt      time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at" struct:"deleted_at"`
}
