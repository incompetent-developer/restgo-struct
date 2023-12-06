package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type TaxInvoice struct {
	ID                   string          `json:"id" struct:"id"`
	PropertyID           string          `json:"property_id" struct:"property_id"`
	Property             *Property       `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	TypeOfMonthID        string          `json:"type_of_month_id" struct:"type_of_month_id" gorm:"foreignKey:type_of_month_id"`
	TypeOfMonth          *TypeOfMonth    `json:"type_of_month,omitempty" struct:"type_of_month,omitempty" gorm:"foreignKey:type_of_month_id"`
	Year                 int64           `json:"year" struct:"year"`
	FileUrl              *string         `json:"file_url,omitempty" struct:"file_url,omitempty"`
	FileStorage          *string         `json:"file_storage,omitempty" struct:"file_storage,omitempty"`
	StatusOfTaxInvoiceID *string         `json:"status_of_tax_invoice_id,omitempty" struct:"status_of_tax_invoice_id,omitempty" gorm:"foreignKey:status_of_tax_invoice_id"`
	StatusOfTaxInvoice   *StatusOfReject `json:"status_of_tax_invoice,omitempty" struct:"status_of_tax_invoice,omitempty" gorm:"foreignKey:status_of_tax_invoice_id"`
	CreatedAt            time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt            *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
