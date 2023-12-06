package cloudsql

import "time"

type UserCreditLog struct {
	ID                   string                 `json:"id" struct:"id" gorm:"primaryKey"`
	StatusOfCreditID     string                 `json:"status_of_credit_id" struct:"status_of_credit_id" gorm:"foreignKey:status_of_credit_id"`
	StatusOfCredit       *StatusOfCredit        `json:"status_of_credit,omitempty" struct:"status_of_credit,omitempty" gorm:"foreignKey:status_of_credit_id"`
	Credit               float64                `json:"credit" struct:"credit"`
	BookingID            *string                `json:"booking_id,omitempty" struct:"booking_id,omitempty"`
	CreatedAt            time.Time              `json:"created_at" struct:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at" struct:"updated_at"`
	UserCreditLogDetails *[]UserCreditLogDetail `json:"user_credit_log_details,omitempty" struct:"user_credit_log_details,omitempty" gorm:"foreignKey:user_credit_log_id"`
}
