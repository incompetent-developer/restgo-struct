package cloudsql

import "time"

// PartnerPaymentTransactionCancelRef structure for mysql
type PartnerPaymentTransactionCancelRef struct {
	OldTransactionID string     `json:"old_transaction_id" structs:"old_transaction_id" gorm:"primary_key;auto_increment:false"`
	NewTransactionID string     `json:"new_transaction_id" structs:"new_transaction_id" gorm:"primary_key;auto_increment:false"`
	CreatedAt        time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID         *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType       *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
