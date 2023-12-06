package cloudsql

import "time"

type PromotionRoomType struct {
	RoomTypeID  string     `json:"room_type_id" struct:"room_type_id" gorm:"foreignKey:room_type_id"`
	RoomType    *RoomType  `json:"room_type,omitempty" struct:"room_type,omitempty" gorm:"foreignKey:room_type_id"`
	PromotionID string     `json:"promotion_id" struct:"promotion_id" gorm:"foreignKey:promotion_id"`
	Promotion   *Promotion `json:"promotion,omitempty" struct:"promotion,omitempty" gorm:"foreignKey:promotion_id"`
	CreatedAt   time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" struct:"updated_at"`
}
