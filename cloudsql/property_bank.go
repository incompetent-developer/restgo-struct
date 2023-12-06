package cloudsql

import "time"

type PropertyBank struct {
	ID          string    `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID  string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property    *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	AccountName string    `json:"account_name" struct:"account_name"`
	BankName    string    `json:"bank_name" struct:"bank_name"`
	BankNumber  string    `json:"bank_number" struct:"bank_number"`
	IsDefault   bool      `json:"is_default" struct:"is_default"`
	CreatedAt   time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" struct:"updated_at"`
}
