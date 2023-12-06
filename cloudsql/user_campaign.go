package cloudsql

import "time"

type UserCampaign struct {
	UserID     string    `json:"user_id" struct:"user_id"`
	User       *User     `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	CampaignID string    `json:"campaign_id" struct:"campaign_id"`
	Campaign   *Campaign `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
