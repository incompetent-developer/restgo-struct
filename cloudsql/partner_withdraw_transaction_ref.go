package cloudsql

import "time"

// PartnerWithdrawTransactionRef structure for mysql
type PartnerWithdrawTransactionRef struct {
	WithdrawTransactionID string                      `json:"withdraw_transaction_id" structs:"withdraw_transaction_id" gorm:"primary_key;auto_increment:false"`
	WithdrawTransaction   *PartnerWithdrawTransaction `json:"withdraw_transaction,omitempty" structs:"withdraw_transaction,omitempty"`
	BookingID             string                      `json:"booking_id" structs:"booking_id" gorm:"primary_key;auto_increment:false"`
	Booking               *Booking                    `json:"bookings,omitempty" structs:"bookings,omitempty"`
	CreatedAt             time.Time                   `json:"created_at" structs:"created_at"`
	UpdatedAt             time.Time                   `json:"updated_at" structs:"updated_at"`
	DeletedAt             *time.Time                  `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID              *string                     `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType            *string                     `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
