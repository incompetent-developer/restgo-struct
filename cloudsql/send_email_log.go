package cloudsql

import "time"

type SendEmailLog struct {
	ID             string    `json:"id" struct:"id" gorm:"primaryKey"`
	SendgridID     string    `json:"sendgrid_id" struct:"sendgrid_id"`
	RecipientEmail string    `json:"recipient_email" struct:"recipient_email"`
	BookingID      *string   `json:"booking_id,omitempty" struct:"booking_id,omitempty"`
	UserID         *string   `json:"user_id,omitempty" struct:"user_id,omitempty"`
	Status         string    `json:"status" struct:"status"`
	Message        string    `json:"message" struct:"message"`
	CreatedAt      time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" struct:"updated_at"`
}
