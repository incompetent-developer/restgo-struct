package cloudsql

import "time"

type FineLog struct {
	FineID                 string                `json:"fine_id" struct:"fine_id" gorm:"foreignKey:fine_id"`
	Fine                   *Fine                 `json:"fine,omitempty" struct:"fine,omitempty" gorm:"foreignKey:fine_id"`
	AccountID              string                `json:"account_id" struct:"account_id"`
	Account                *Account              `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	StatusFineOfTransferID string                `json:"status_fine_of_transfer_id" struct:"status_fine_of_transfer_id" gorm:"foreignKey:status_fine_of_transfer_id"`
	StatusFineOfTransfer   *StatusFineOfTransfer `json:"status_fine_of_transfer,omitempty" struct:"status_fine_of_transfer,omitempty" gorm:"foreignKey:status_fine_of_transfer_id"`
	Description            *string               `json:"description,omitempty" struct:"description,omitempty"`
	CreatedAt              time.Time             `json:"created_at" struct:"created_at"`
	UpdatedAt              time.Time             `json:"updated_at" struct:"updated_at"`
}
