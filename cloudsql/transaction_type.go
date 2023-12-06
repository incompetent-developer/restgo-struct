package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType struct {
	ID         string          `json:"id" struct:"id" gorm:"primaryKey"`
	Name       string          `json:"name" struct:"name"`
	CreatedAt  time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	EditorID   *string         `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType *string         `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
