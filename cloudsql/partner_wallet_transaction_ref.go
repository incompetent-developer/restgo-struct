package cloudsql

import "time"

// PartnerWalletTransactionRef structure for mysql
type PartnerWalletTransactionRef struct {
	WalletTransactionID string                    `json:"wallet_transaction_id" structs:"wallet_transaction_id" gorm:"primary_key;auto_increment:false"`
	WalletTransaction   *PartnerWalletTransaction `json:"wallet_transactions,omitempty" structs:"wallet_transactions,omitempty"`
	BookingID           string                    `json:"booking_id" structs:"booking_id" gorm:"primary_key;auto_increment:false"`
	Booking             *Booking                  `json:"bookings,omitempty" structs:"bookings,omitempty"`
	CreatedAt           time.Time                 `json:"created_at" structs:"created_at"`
	UpdatedAt           time.Time                 `json:"updated_at" structs:"updated_at"`
	DeletedAt           *time.Time                `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID            *string                   `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType          *string                   `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
