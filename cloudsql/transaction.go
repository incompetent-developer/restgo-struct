package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                        string                      `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID                string                      `json:"property_id" struct:"property_id"`
	Property                  *Property                   `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	PrepayDate                *time.Time                  `json:"prepay_date" struct:"prepay_date"`
	CheckoutDate              *time.Time                  `json:"checkout_date" struct:"checkout_date"`
	TotalBasePrice            float64                     `json:"total_base_price" struct:"total_base_price"`
	TotalAddon                *float64                    `json:"total_addon" struct:"total_addon"`
	TotalTax3                 *float64                    `json:"total_tax3" struct:"total_tax3"`
	TotalOffer                *float64                    `json:"total_offer" struct:"total_offer"`
	TotalPet                  *float64                    `json:"total_pet" struct:"total_pet"`
	TotalExtraBed             *float64                    `json:"total_extra_bed" struct:"total_extra_bed"`
	TotalDiscount             *float64                    `json:"total_discount" struct:"total_discount"`
	TotalCancelCharge         *float64                    `json:"total_cancel_charge" struct:"total_cancel_charge"`
	TotalVat                  *float64                    `json:"total_vat" struct:"total_vat"`
	TotalFee                  *float64                    `json:"total_fee" struct:"total_fee"`
	TotalPrice                float64                     `json:"total_price" struct:"total_price"`
	PartnerAccountID          *string                     `json:"partner_account_id" struct:"partner_account_id"`
	PartnerAccount            *Account                    `json:"partner_account,omitempty" struct:"partner_account,omitempty" gorm:"foreignKey:partner_account_id"`
	PartnerTypeOfProcessID    *string                     `json:"partner_type_of_process_id" struct:"partner_type_of_process_id"` //ยื่นอนุมัติ, ยื่นปฏิเสธ
	PartnerTypeOfProcess      *PartnerTypeOfProcess       `json:"partner_type_of_process,omitempty" struct:"partner_type_of_process,omitempty" gorm:"foreignKey:partner_type_of_process_id"`
	StatusOfTransactionID     *string                     `json:"status_of_transaction_id" struct:"status_of_transaction_id"` //รออนุมัติ, อนุมัติแล้ว, ยกเลิกโดยระบบ
	StatusOfTransaction       *StatusOfTransaction        `json:"status_of_transaction,omitempty" struct:"status_of_transaction,omitempty" gorm:"foreignKey:status_of_transaction_id"`
	PartnerProcessDate        *time.Time                  `json:"partner_process_date" struct:"partner_process_date"` //วันยื่นอนุมัติ/วันยื่นปฏิเสธ
	PartnerRejectDate         *time.Time                  `json:"partner_reject_date" struct:"partner_reject_date"`
	PartnerRejectDescription  *string                     `json:"partner_reject_description" struct:"partner_reject_description"` //กรณียื่นปฏิเสธ : คําอธิบายการยื่นปฏิเสธ
	StatusOfRejectID          *string                     `json:"status_of_reject_id" struct:"status_of_reject_id"`               //รอการตรวจสอบ, อยู่ระหว่างการตรวจสอบ, ตรวจสอบเรียบร้อยแล้ว
	StatusOfReject            *StatusOfReject             `json:"status_of_reject,omitempty" struct:"status_of_reject,omitempty" gorm:"foreignKey:status_of_reject_id"`
	AeTypeOfProcessID         *string                     `json:"ae_type_of_process_id" struct:"ae_type_of_process_id"` //ข้อมูลถูกต้อง, ข้อมูลไม่ถูกต้อง
	AeTypeOfProcess           *AeTypeOfProcess            `json:"ae_type_of_process,omitempty" struct:"ae_type_of_process,omitempty" gorm:"foreignKey:ae_type_of_process_id"`
	AcTypeOfProcessID         *string                     `json:"type_of_process_id" struct:"type_of_process_id"` //รอการตรวจสอบ, รออนุมัติ, รายการไม่ถูกต้อง, อนุมัติแล้ว
	AcTypeOfProcess           *AcTypeOfProcess            `json:"ac_type_of_process,omitempty" struct:"ac_type_of_process,omitempty" gorm:"foreignKey:ac_type_of_process_id"`
	ProcessDate               *time.Time                  `json:"process_date" struct:"process_date"`                         // วันยื่นจาก partner
	ReferenceTransactionID    *string                     `json:"reference_transaction_id" struct:"reference_transaction_id"` // อ้างอิง
	TransactionTypeID         string                      `json:"transaction_type_id" struct:"transaction_type_id"`           // 'checkout', 'prepay', 'pay-at', 'fine'
	InvoiceUrl                *string                     `json:"invoice_url" struct:"invoice_url"`
	InvoiceStorage            *string                     `json:"invoice_storage" struct:"invoice_storage"`
	TaxInvoiceUrl             *string                     `json:"tax_invoice_url" struct:"tax_invoice_url"`
	TaxInvoiceStorage         *string                     `json:"tax_invoice_storage" struct:"tax_invoice_storage"`
	IsOverdue                 bool                        `json:"is_overdue" struct:"is_overdue"` // default 7 วัน แก้ได้
	CreatedAt                 time.Time                   `json:"created_at" struct:"created_at"`
	UpdatedAt                 time.Time                   `json:"updated_at" struct:"updated_at"`
	DeletedAt                 *gorm.DeletedAt             `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	TransactionImages         *[]TransactionImage         `json:"transaction_images,omitempty" struct:"transaction_images,omitempty" gorm:"foreignKey:transaction_id"`
	TransactionBookings       *[]TransactionBooking       `json:"transaction_bookings,omitempty" struct:"transaction_bookings,omitempty" gorm:"foreignKey:transaction_id"`
	TransactionBookingDetails *[]TransactionBookingDetail `json:"transaction_booking_details,omitempty" struct:"transaction_booking_details,omitempty" gorm:"foreignKey:transaction_id"`
	TransactionLogs           *[]TransactionLog           `json:"transaction_logs,omitempty" struct:"transaction_logs,omitempty" gorm:"foreignKey:transaction_id"`
}
