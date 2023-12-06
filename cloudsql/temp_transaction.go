package cloudsql

import "time"

type TempTransaction struct {
	ID                      string                    `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID              string                    `json:"property_id" struct:"property_id"`
	Property                *Property                 `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	PrepayDate              *time.Time                `json:"prepay_date,omitempty" struct:"prepay_date,omitempty"`
	CheckoutDate            *time.Time                `json:"checkout_date,omitempty" struct:"checkout_date,omitempty"`
	TotalBasePrice          float64                   `json:"total_base_price" struct:"total_base_price"`
	TotalAddon              *float64                  `json:"total_addon" struct:"total_addon"`
	TotalTax3               *float64                  `json:"total_tax3" struct:"total_tax3"`
	TotalOffer              *float64                  `json:"total_offer,omitempty" struct:"total_offer,omitempty"`
	TotalPet                *float64                  `json:"total_pet,omitempty" struct:"total_pet,omitempty"`
	TotalExtraBed           *float64                  `json:"total_extra_bed,omitempty" struct:"total_extra_bed,omitempty"`
	TotalDiscount           *float64                  `json:"total_discount,omitempty" struct:"total_discount,omitempty"`
	TotalCancelCharge       *float64                  `json:"total_cancel_charge,omitempty" struct:"total_cancel_charge,omitempty"`
	TotalVat                *float64                  `json:"total_vat" struct:"total_vat"`
	TotalFee                *float64                  `json:"total_fee,omitempty" struct:"total_fee,omitempty"`
	TotalPrice              float64                   `json:"total_price" struct:"total_price"`
	StatusOfTransactionID   *string                   `json:"status_of_transaction_id,omitempty" struct:"status_of_transaction_id,omitempty"` //รอการตรวจสอบ, ข้อมูลถูกต้อง, ข้อมูลไม่ถูกต้อง
	StatusOfTransaction     *StatusOfTransaction      `json:"status_of_transaction,omitempty" struct:"status_of_transaction,omitempty" gorm:"foreignKey:status_of_transaction_id"`
	ReferenceTransactionID  *string                   `json:"reference_transaction_id,omitempty" struct:"reference_transaction_id,omitempty"` // อ้างอิง
	TransactionTypeID       string                    `json:"transaction_type_id" struct:"transaction_type_id"`                               // 'transaction', 'transaction-จ่าย ณ ที่พัก', 'fine'
	IsOverdue               bool                      `json:"is_overdue" struct:"is_overdue"`
	CreatedAt               time.Time                 `json:"created_at" struct:"created_at"`
	UpdatedAt               time.Time                 `json:"updated_at" struct:"updated_at"`
	TempTransactionBookings *[]TempTransactionBooking `json:"temp_transaction_bookings,omitempty" struct:"temp_transaction_bookings,omitempty" gorm:"foreignKey:temp_transaction_id"`
	TempTransactionLogs     *[]TempTransactionLog     `json:"temp_transaction_logs,omitempty" struct:"temp_transaction_logs,omitempty" gorm:"foreignKey:temp_transaction_id"`
}
