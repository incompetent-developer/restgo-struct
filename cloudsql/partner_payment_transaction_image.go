package cloudsql

import "time"

// PartnerPaymentTransactionImage structure for mysql
type PartnerPaymentTransactionImage struct {
	ID                   uint64                     `json:"id" structs:"id"`
	PaymentTransactionID string                     `json:"payment_transaction_id" structs:"payment_transaction_id"`
	PaymentTransaction   *PartnerPaymentTransaction `json:"payment_transaction,omitempty" structs:"payment_transaction,omitempty"`
	ImagePath            string                     `json:"image_path" structs:"image_path"`
	Date                 string                     `json:"date" structs:"date"`
	Time                 string                     `json:"time" structs:"time"`
	Amount               float64                    `json:"amount" structs:"amount"`
	RestgoBankID         string                     `json:"restgo_bank_id" structs:"restgo_bank_id"`
	RestgoBank           *RestgoBank                `json:"restgo_bank,omitempty" structs:"restgo_bank,omitempty"`
	CreatedAt            time.Time                  `json:"created_at" structs:"created_at"`
	UpdatedAt            time.Time                  `json:"updated_at" structs:"updated_at"`
	DeletedAt            *time.Time                 `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID             *string                    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType           *string                    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
