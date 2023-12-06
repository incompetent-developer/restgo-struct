package cloudsql

import "time"

type CampaignTimeSale struct {
	CampaignID string    `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign   *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	TimeStart  string    `json:"time_start" struct:"time_start"`
	TimeEnd    string    `json:"time_end" struct:"time_end"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
