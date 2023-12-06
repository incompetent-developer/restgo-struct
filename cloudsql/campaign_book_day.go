package cloudsql

import "time"

type CampaignBookDay struct {
	CampaignID  string     `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign    *Campaign  `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	DayOfWeekID string     `json:"day_of_week_id" struct:"day_of_week_id" gorm:"foreignKey:day_of_week_id"`
	DayOfWeek   *DayOfWeek `json:"day_of_week,omitempty" struct:"day_of_week,omitempty" gorm:"foreignKey:day_of_week_id"`
	CreatedAt   time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" struct:"updated_at"`
}
