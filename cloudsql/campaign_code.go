package cloudsql

import "time"

type CampaignCode struct {
	CampaignID string    `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign   *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	Code       string    `json:"code" struct:"code"`
	IsMain     bool      `json:"is_main" struct:"is_main"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
