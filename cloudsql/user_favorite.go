package cloudsql

import "time"

type UserFavorite struct {
	UserID     string    `json:"user_id" struct:"user_id" gorm:"foreignKey:user_id"`
	User       *User     `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	PropertyID string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property   *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
