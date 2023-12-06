package cloudsql

import "time"

type AccountPermission struct {
	AccountID    string      `json:"account_id" struct:"account_id" gorm:"foreignKey:account_id"`
	Account      *Account    `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	PermissionID string      `json:"permission_id" struct:"permission_id" gorm:"foreignKey:permission_id"`
	Permission   *Permission `json:"permission,omitempty" struct:"permission,omitempty" gorm:"foreignKey:permission_id"`
	IsEnabled    bool        `json:"is_enabled" struct:"is_enabled"`
	CreatedAt    time.Time   `json:"created_at" struct:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" struct:"updated_at"`
}
