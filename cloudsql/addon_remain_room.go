package cloudsql

import "time"

type AddonRemainRoom struct {
	ID                string              `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID        *string             `json:"property_id,omitempty" struct:"property_id,omitempty" gorm:"foreignKey:property_id,omitempty"`
	Property          *Property           `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	IsDefault         bool                `json:"is_default" struct:"is_default"`
	CreatedAt         time.Time           `json:"created_at" struct:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at" struct:"updated_at"`
	AddonRemainRanges *[]AddonRemainRange `json:"addon_remain_ranges,omitempty" struct:"addon_remain_ranges,omitempty" gorm:"foreignKey:addon_remain_room_id"`
}
