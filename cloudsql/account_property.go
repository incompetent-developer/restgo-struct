package cloudsql

import "time"

type AccountProperty struct {
	AccountID  string    `json:"account_id" struct:"account_id" gorm:"foreignKey:account_id"`
	Account    *Account  `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	PropertyID string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property   *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	IsEnabled  bool      `json:"is_enabled" struct:"is_enabled"`
	CreatedAt  time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" struct:"updated_at"`
}
