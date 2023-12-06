package cloudsql

import "time"

type TempTransactionLog struct {
	TempTransactionID string           `json:"temp_transaction_id" struct:"temp_transaction_id" gorm:"primaryKey"`
	TempTransaction   *TempTransaction `json:"temp_transaction,omitempty" struct:"temp_transaction,omitempty" gorm:"foreignKey:temp_transaction_id"`
	AccountID         string           `json:"account_id" struct:"account_id"`
	Account           *Account         `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	TypeOfProcessID   string           `json:"type_of_process_id" struct:"type_of_process_id"` // ข้อมูลถูกต้อง, ข้อมูลไม่ถูกต้อง
	TypeOfProcess     *AeTypeOfProcess `json:"type_of_process,omitempty" struct:"type_of_process,omitempty" gorm:"foreignKey:type_of_process_id"`
	Description       *string          `json:"description,omitempty" struct:"description,omitempty"` // อธิบายเวลายื่นข้อมูลถูกต้อง(AE), ข้อมูลไม่ถูกต้อง (AC)
	CreatedAt         time.Time        `json:"created_at" struct:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at" struct:"updated_at"`
}
