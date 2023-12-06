package cloudsql

import "time"

type CampaignPropertyProvince struct {
	CampaignID string    `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign   *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	Province   string    `json:"province" struct:"province" gorm:"foreignKey:province"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
