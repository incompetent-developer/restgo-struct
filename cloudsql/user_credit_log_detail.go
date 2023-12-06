package cloudsql

import "time"

type UserCreditLogDetail struct {
	UserCreditID    string         `json:"user_credit_id" struct:"user_credit_id" gorm:"foreignKey:user_credit_id"`
	UserCredit      *UserCredit    `json:"user_credit,omitempty" struct:"user_credit,omitempty" gorm:"foreignKey:user_credit_id"`
	UserCreditLogID string         `json:"user_credit_log_id" struct:"user_credit_log_id" gorm:"foreignKey:user_credit_log_id"`
	UserCreditLog   *UserCreditLog `json:"user_credit_log,omitempty" struct:"user_credit_log,omitempty" gorm:"foreignKey:user_credit_log_id"`
	Credit          float64        `json:"credit" struct:"credit"`
	CreatedAt       time.Time      `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" struct:"updated_at"`
}
