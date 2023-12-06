package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type ReasonOfCancel struct {
	ID             string          `json:"id" struct:"id" gorm:"primaryKey"`
	Name           string          `json:"name" struct:"name"`
	TypeOfReasonID string          `json:"type_of_reason_id" struct:"type_of_reason_id" gorm:"foreignKey:type_of_reason_id"`
	TypeOfReason   *TypeOfReason   `json:"type_of_reason,omitempty" struct:"type_of_reason,omitempty" gorm:"foreignKey:type_of_reason_id"`
	IsSuspended    bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt      time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
