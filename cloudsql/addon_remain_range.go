package cloudsql

import "time"

type AddonRemainRange struct {
	AddonRemainRoomID string           `json:"addon_remain_room_id" struct:"addon_remain_room_id" gorm:"foreignKey:addon_remain_room_id"`
	AddonRemainRoom   *AddonRemainRoom `json:"addon_remain_room,omitempty" struct:"addon_remain_room,omitempty" gorm:"foreignKey:addon_remain_room_id"`
	RemainRangeID     string           `json:"remain_range_id,omitempty" struct:"remain_range_id,omitempty" gorm:"foreignKey:remain_range_id"`
	RemainRange       *RemainRange     `json:"remain_range,omitempty" struct:"remain_range,omitempty" gorm:"foreignKey:remain_range_id"`
	Compensation      int64            `json:"compensation" struct:"compensation"`
	CompensationMin   int64            `json:"compensation_min" struct:"compensation_min"`
	CompensationMax   int64            `json:"compensation_max" struct:"compensation_max"`
	CreatedAt         time.Time        `json:"created_at" struct:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at" struct:"updated_at"`
}
