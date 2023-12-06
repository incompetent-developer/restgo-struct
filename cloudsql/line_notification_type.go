package cloudsql

import "time"

// LineNotificationType structure for mysql
type LineNotificationType struct {
	ID          string     `json:"id" structs:"id" gorm:"primary_key"`
	Description string     `json:"description" structs:"description"`
	IsEnabled   bool       `json:"is_enabled" structs:"is_enabled"`
	CreatedAt   time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID    *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType  *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
