package cloudsql

import "time"

type UserCredit struct {
	ID                    string               `json:"id" struct:"id" gorm:"primaryKey"`
	UserID                string               `json:"user_id" struct:"user_id"  gorm:"foreignKey:user_id"`
	User                  *User                `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	Credit                float64              `json:"credit" struct:"credit"`
	StatusOfCreditID      string               `json:"status_of_credit_id" struct:"status_of_credit_id"  gorm:"foreignKey:status_of_credit_id"`
	StatusOfCredit        *StatusOfCredit      `json:"status_of_credit,omitempty" struct:"status_of_credit,omitempty" gorm:"foreignKey:status_of_credit_id"`
	TypeOfCreditReceiveID *string              `json:"type_of_credit_receive_id,omitempty" struct:"type_of_credit_receive_id,omitempty"  gorm:"foreignKey:type_of_credit_receive_id"`
	TypeOfCreditReceive   *TypeOfCreditReceive `json:"type_of_credit_receive,omitempty" struct:"type_of_credit_receive,omitempty" gorm:"foreignKey:type_of_credit_receive_id"`
	ExpiredTime           *time.Time           `json:"expired_time,omitempty" struct:"expired_time,omitempty"`
	CreatedAt             time.Time            `json:"created_at" struct:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at" struct:"updated_at"`
}
