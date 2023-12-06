package cloudsql

import "time"

type CampaignPropertyRegion struct {
	CampaignID string    `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign   *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	Region     string    `json:"region" struct:"region" gorm:"foreignKey:region"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
