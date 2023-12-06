package cloudsql

import "time"

type CampaignPropertySubDistrict struct {
	CampaignID  string    `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign    *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	SubDistrict string    `json:"sub_district" struct:"sub_district" gorm:"foreignKey:sub_district"`
	CreatedAt   time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" struct:"updated_at"`
}
