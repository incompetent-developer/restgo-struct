package cloudsql

import "time"

type BookingCampaign struct {
	BookingID     string    `json:"booking_id" struct:"booking_id"`
	Booking       *Booking  `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	CampaignID    string    `json:"campaign_id" struct:"campaign_id"`
	Campaign      *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPrice float64   `json:"campaign_price" struct:"campaign_price"`
	CreatedAt     time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" struct:"updated_at"`
}
