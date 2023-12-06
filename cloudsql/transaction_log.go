package cloudsql

import "time"

type TransactionLog struct {
	TransactionID          string               `json:"transaction_id" struct:"transaction_id" gorm:"primaryKey"`
	Transaction            *Transaction         `json:"transaction,omitempty" struct:"transaction,omitempty" gorm:"foreignKey:transaction_id"`
	AccountID              string               `json:"account_id" struct:"account_id"`
	Account                *Account             `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	AeTypeOfProcessID      *string              `json:"ae_type_of_process_id" struct:"ae_type_of_process_id"` //ข้อมูลถูกต้อง, ข้อมูลไม่ถูกต้อง
	AeTypeOfProcess        *AeTypeOfProcess     `json:"ae_type_of_process,omitempty" struct:"ae_type_of_process,omitempty" gorm:"foreignKey:ae_type_of_process_id"`
	AcTypeOfProcessID      *string              `json:"ac_type_of_process_id" struct:"ac_type_of_process_id"` //รอการตรวจสอบ, รออนุมัติ, รายการไม่ถูกต้อง, อนุมัติแล้ว
	AcTypeOfProcess        *AcTypeOfProcess     `json:"ac_type_of_process,omitempty" struct:"ac_type_of_process,omitempty" gorm:"foreignKey:ac_type_of_process_id"`
	PartnerTypeOfProcessID string               `json:"partner_type_of_process_id" struct:"partner_type_of_process_id"` //ยื่นอนุมัติ, ยื่นปฏิเสธ
	PartnerTypeOfProcess   PartnerTypeOfProcess `json:"partner_type_of_process" struct:"partner_type_of_process" gorm:"foreignKey:partner_type_of_process_id"`
	Description            *string              `json:"description" struct:"description"`
	Note                   *string              `json:"note" struct:"note"` //ผลการตรวจสอบ : หมายเหตุ ใส่เลขธุรกรรมใบใหม่
	CreatedAt              time.Time            `json:"created_at" struct:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at" struct:"updated_at"`
}
