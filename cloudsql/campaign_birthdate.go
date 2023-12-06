package cloudsql

import "time"

type CampaignBirthdate struct {
	CampaignID     string     `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign       *Campaign  `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	BirthdateStart time.Time  `json:"birthdate_start" struct:"birthdate_start"`
	BirthdateEnd   *time.Time `json:"birthdate_end,omitempty" struct:"birthdate_end,omitempty"`
	CreatedAt      time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" struct:"updated_at"`
}
