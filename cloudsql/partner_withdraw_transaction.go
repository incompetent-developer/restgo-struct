package cloudsql

import (
	"time"
)

//PartnerWithdrawTransaction structure
type PartnerWithdrawTransaction struct {
	ID                                string                             `json:"id" structs:"id"`
	WalletTransactionID               string                             `json:"wallet_transaction_id" structs:"wallet_transaction_id"`
	PropertyID                        string                             `json:"property_id" structs:"property_id"`
	PropertyBankID                    string                             `json:"property_bank_id" structs:"property_bank_id"`
	Amount                            float64                            `json:"amount" structs:"amount"`
	BasePrice                         *float64                           `json:"base_price,omitempty" structs:"base_price,omitempty"`
	FeeRate                           *float64                           `json:"fee_rate,omitempty" structs:"fee_rate,omitempty"`
	VatRate                           *float64                           `json:"vat_rate,omitempty" structs:"vat_rate,omitempty"`
	StatusID                          string                             `json:"status_id" structs:"status_id"`
	AEStatusID                        string                             `json:"ae_status_id" structs:"ae_status_id"`
	ACStatusID                        string                             `json:"ac_status_id" structs:"ac_status_id"`
	PaymentPlatformID                 *string                            `json:"payment_platform_id,omitempty" structs:"payment_platform_id,omitempty"`
	PaymentSubplatformID              *string                            `json:"payment_subplatform_id,omitempty" structs:"payment_subplatform_id,omitempty"`
	TransactionCreatedTime            time.Time                          `json:"transaction_created_time" structs:"transaction_created_time"`
	TransactionSuccessAt              *time.Time                         `json:"transaction_success_at,omitempty" structs:"transaction_success_at,omitempty"`
	Description                       *string                            `json:"description,omitempty" structs:"description,omitempty"`
	PNUpdatedAt                       *time.Time                         `json:"pn_updated_at,omitempty" structs:"pn_updated_at,omitempty"`
	PNReportUpdatedAt                 *time.Time                         `json:"pn_report_updated_at,omitempty" structs:"pn_report_updated_at,omitempty"`
	AEInspectorID                     *string                            `json:"ae_inspector_id,omitempty" structs:"ae_inspector_id,omitempty"`
	AEUpdatedAt                       *time.Time                         `json:"ae_updated_at,omitempty" structs:"ae_updated_at,omitempty"`
	ACInspectorID                     *string                            `json:"ac_inspector_id,omitempty" structs:"ac_inspector_id,omitempty"`
	ACUpdatedAt                       *time.Time                         `json:"ac_updated_at,omitempty" structs:"ac_updated_at,omitempty"`
	ImagePath                         *string                            `json:"image_path,omitempty" structs:"image_path,omitempty"`
	CreatorID                         *string                            `json:"creator_id,omitempty" structs:"creator_id,omitempty"`
	CreatorType                       *string                            `json:"creator_type,omitempty" structs:"creator_type,omitempty"`
	CreatedAt                         time.Time                          `json:"created_at" structs:"created_at"`
	UpdatedAt                         time.Time                          `json:"updated_at" structs:"updated_at"`
	DeletedAt                         *time.Time                         `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID                          *string                            `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType                        *string                            `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
	PartnerWithdrawTransactionReceipt *PartnerWithdrawTransactionReceipt `json:"partner_withdraw_transaction_receipt,omitempty" struct:"partner_withdraw_transaction_receipt,omitempty" gorm:"foreignKey:withdraw_transaction_id"`
}
