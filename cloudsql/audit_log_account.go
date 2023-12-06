package cloudsql

import "time"

type AuditLogAccount struct {
	AccountID          string    `json:"account_id" struct:"account_id" gorm:"foreignKey:account_id"`
	Account            *Account  `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	ResourceName       string    `json:"resource_name" struct:"resource_name"`
	Type               string    `json:"type" struct:"type"`
	Action             string    `json:"action" struct:"action"`
	PerformedBy        string    `json:"performed_by" struct:"performed_by"`
	Parameters         string    `json:"parameters" struct:"parameters"`
	PreviousParameters string    `json:"previous_parameters" struct:"previous_parameters"`
	CreatedAt          time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" struct:"updated_at"`
}
