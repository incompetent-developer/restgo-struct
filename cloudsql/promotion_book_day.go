package cloudsql

import "time"

type PromotionBookDay struct {
	PromotionID string     `json:"promotion_id" struct:"promotion_id" gorm:"foreignKey:promotion_id"`
	Promotion   *Promotion `json:"promotion,omitempty" struct:"promotion,omitempty" gorm:"foreignKey:promotion_id"`
	DayOfWeekID string     `json:"day_of_week_id" struct:"day_of_week_id" gorm:"foreignKey:day_of_week_id"`
	DayOfWeek   *DayOfWeek `json:"day_of_week,omitempty" struct:"day_of_week,omitempty" gorm:"foreignKey:day_of_week_id"`
	CreatedAt   time.Time  `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" struct:"updated_at"`
}
