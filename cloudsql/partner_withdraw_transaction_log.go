package cloudsql

import (
	"time"
)

//PartnerWithdrawTransaction structure
type PartnerWithdrawTransactionLog struct {
	ID                    uint64     `json:"id" structs:"id"`
	WithdrawTransactionID string     `json:"withdraw_transaction_id" structs:"withdraw_transaction_id"`
	Amount                float64    `json:"amount" structs:"amount"`
	BasePrice             *float64   `json:"base_price,omitempty" structs:"base_price,omitempty"`
	FeeRate               *float64   `json:"fee_rate,omitempty" structs:"fee_rate,omitempty"`
	VatRate               *float64   `json:"vat_rate,omitempty" structs:"vat_rate,omitempty"`
	StatusID              string     `json:"status_id" structs:"status_id"`
	AEStatusID            string     `json:"ae_status_id" structs:"ae_status_id"`
	ACStatusID            string     `json:"ac_status_id" structs:"ac_status_id"`
	Description           *string    `json:"description,omitempty" structs:"description,omitempty"`
	PNUpdatedAt           *time.Time `json:"pn_updated_at,omitempty" structs:"pn_updated_at,omitempty"`
	PNReportUpdatedAt     *time.Time `json:"pn_report_updated_at,omitempty" structs:"pn_report_updated_at,omitempty"`
	AEInspectorID         *string    `json:"ae_inspector_id,omitempty" structs:"ae_inspector_id,omitempty"`
	AEUpdatedAt           *time.Time `json:"ae_updated_at,omitempty" structs:"ae_updated_at,omitempty"`
	ACInspectorID         *string    `json:"ac_inspector_id,omitempty" structs:"ac_inspector_id,omitempty"`
	ACUpdatedAt           *time.Time `json:"ac_updated_at,omitempty" structs:"ac_updated_at,omitempty"`
	CreatedAt             time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID              *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType            *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
