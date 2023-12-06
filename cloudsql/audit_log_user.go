package cloudsql

import "time"

type AuditLogUser struct {
	UserID    string    `json:"user_id" struct:"user_id" gorm:"foreignKey:user_id"`
	User      *User     `json:"user" struct:"user" gorm:"foreignKey:user_id"`
	Action    string    `json:"action" struct:"action"`
	Function  string    `json:"function" struct:"function"`
	TableName string    `json:"table_name" struct:"table_name"`
	Code      int64     `json:"code" struct:"code"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
