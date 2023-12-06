package cloudsql

import "time"

type CampaignMonth struct {
	CampaignID    string       `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign      *Campaign    `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	TypeOfMonthID string       `json:"type_of_month_id" struct:"type_of_month_id" gorm:"foreignKey:type_of_month_id"`
	TypeOfMonth   *TypeOfMonth `json:"type_of_month,omitempty" struct:"type_of_month,omitempty" gorm:"foreignKey:type_of_month_id"`
	CreatedAt     time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" struct:"updated_at"`
}
