package cloudsql

import (
	"time"

	"github.com/incompetent-developer/restgo-struct/status"
	"gorm.io/gorm"
)

//PartnerWalletTransaction structure
type PartnerWalletTransaction struct {
	ID                     string     `json:"id" structs:"id"`
	PropertyID             string     `json:"property_id" structs:"property_id"`
	BookingID              *string    `json:"booking_id,omitempty" structs:"booking_id,omitempty"`
	Amount                 float64    `json:"amount" structs:"amount"`
	BasePrice              *float64   `json:"base_price,omitempty" structs:"base_price,omitempty"`
	FeeRate                *float64   `json:"fee_rate,omitempty" structs:"fee_rate,omitempty"`
	VatRate                *float64   `json:"vat_rate,omitempty" structs:"vat_rate,omitempty"`
	MarkupRate             *float64   `json:"markup_rate,omitempty" structs:"markup_rate,omitempty"`
	WithholdingRate        *float64   `json:"withholding_rate,omitempty" structs:"withholding_rate,omitempty"`
	TransactionTypeID      string     `json:"transaction_type_id" structs:"transaction_type_id"`
	TransactionCreatedTime time.Time  `json:"transaction_created_time" structs:"transaction_created_time"`
	TransactionSuccessAt   *time.Time `json:"transaction_success_at,omitempty" structs:"transaction_success_at,omitempty"`
	StatusID               string     `json:"status_id" structs:"status_id"`
	BookingPaymentTypeID   *string    `json:"booking_payment_type_id,omitempty" structs:"booking_payment_type_id,omitempty"`
	PaymentPlatformID      *string    `json:"payment_platform_id,omitempty" structs:"payment_platform_id,omitempty"`
	PaymentSubplatformID   *string    `json:"payment_subplatform_id,omitempty" structs:"payment_subplatform_id,omitempty"`
	IsTXSuccessHooked      bool       `json:"-" structs:"-"`
	CreatedAt              time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID               *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType             *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}

// BeforeSave action before save wallet_transactions
func (wt *PartnerWalletTransaction) BeforeSave(tx *gorm.DB) (err error) {
	/*
		Cached
	*/
	if wt.StatusID == status.Success && !wt.IsTXSuccessHooked {

		/*
			Ensure success at
		*/
		tn := time.Now()
		wt.TransactionSuccessAt = &tn

		/*
			Prepare save
		*/
		wt.IsTXSuccessHooked = true

	}

	return nil
}
