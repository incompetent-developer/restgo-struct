package cloudsql

import "time"

type TempBookingCampaign struct {
	TempBookingID string       `json:"temp_booking_id" struct:"temp_booking_id"`
	TempBooking   *TempBooking `json:"temp_booking,omitempty" struct:"temp_booking,omitempty" gorm:"foreignKey:temp_booking_id"`
	CampaignID    string       `json:"campaign_id" struct:"campaign_id"`
	Campaign      *Campaign    `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPrice float64      `json:"campaign_price" struct:"campaign_price"`
	CreatedAt     time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" struct:"updated_at"`
}
