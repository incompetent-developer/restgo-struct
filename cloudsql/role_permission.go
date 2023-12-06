package cloudsql

import "time"

type RolePermission struct {
	RoleID       string      `json:"role_id" struct:"role_id" gorm:"foreignKey:role_id"`
	Role         *Role       `json:"role,omitempty" struct:"role,omitempty" gorm:"foreignKey:role_id"`
	PermissionID string      `json:"permission_id" struct:"permission_id" gorm:"foreignKey:permission_id"`
	Permission   *Permission `json:"permission,omitempty" struct:"permission,omitempty" gorm:"foreignKey:permission_id"`
	IsEnabled    bool        `json:"is_enabled" struct:"is_enabled"`
	CreatedAt    time.Time   `json:"created_at" struct:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" struct:"updated_at"`
}
