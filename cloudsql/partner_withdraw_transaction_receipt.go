package cloudsql

import "time"

// PartnerWithdrawTransactionReceipt structure for mysql
type PartnerWithdrawTransactionReceipt struct {
	ID                         string                      `json:"id" struct:"id" gorm:"primaryKey"`
	WithdrawTransactionID      string                      `json:"withdraw_transaction_id" struct:"withdraw_transaction_id" gorm:"foreignKey:withdraw_transaction_id"`
	PartnerWithdrawTransaction *PartnerWithdrawTransaction `json:"partner_withdraw_transaction,omitempty" struct:"partner_withdraw_transaction,omitempty" gorm:"foreignKey:withdraw_transaction_id"`
	ImagePath                  string                      `json:"image_path" struct:"image_path"`
	AEStatusID                 string                      `json:"ae_status_id" structs:"ae_status_id"`
	CreatedAt                  time.Time                   `json:"created_at" structs:"created_at"`
	UpdatedAt                  time.Time                   `json:"updated_at" structs:"updated_at"`
	DeletedAt                  *time.Time                  `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID                   *string                     `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType                 *string                     `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
