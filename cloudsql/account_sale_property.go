package cloudsql

import "time"

type AccountSaleProperty struct {
	AccountSaleID string       `json:"account_sale_id" struct:"account_sale_id" gorm:"foreignKey:account_sale_id"`
	AccountSale   *AccountSale `json:"account_sale,omitempty" struct:"account_sale,omitempty" gorm:"foreignKey:account_sale_id"`
	PropertyID    string       `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property      *Property    `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	CreatedAt     time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" struct:"updated_at"`
}
