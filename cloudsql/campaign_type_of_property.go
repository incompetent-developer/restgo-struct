package cloudsql

import "time"

type CampaignTypeOfProperty struct {
	CampaignID       string          `json:"campaign_id" struct:"campaign_id" gorm:"foreignKey:campaign_id"`
	Campaign         *Campaign       `json:"campaign,omitempty" struct:"campaign,omitempty" gorm:"foreignKey:campaign_id"`
	TypeOfPropertyID string          `json:"type_of_property_id" struct:"type_of_property_id" gorm:"foreignKey:type_of_property_id"`
	TypeOfProperty   *TypeOfProperty `json:"type_of_property,omitempty" struct:"type_of_property,omitempty" gorm:"foreignKey:type_of_property_id"`
	CreatedAt        time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" struct:"updated_at"`
}
