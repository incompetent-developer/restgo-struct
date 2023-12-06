package cloudsql

import "time"

type HistoryLogin struct {
	AccountID string    `json:"account_id" struct:"account_id" gorm:"foreignKey:account_id"`
	Account   *Account  `json:"account" struct:"account" gorm:"foreignKey:account_id"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
