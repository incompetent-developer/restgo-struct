package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID                 string               `json:"id" struct:"id" gorm:"primaryKey"`
	Firstname          *string              `json:"firstname" struct:"firstname"`
	Lastname           *string              `json:"lastname" struct:"lastname"`
	Password           string               `json:"-" struct:"password"`
	RoleID             string               `json:"role_id" struct:"role_id" gorm:"foreignKey:role_id"`
	Role               *Role                `json:"role,omitempty" struct:"role,omitempty" gorm:"foreignKey:role_id"`
	Email              *string              `json:"email" struct:"email"`
	TypeOfPhoneID      *string              `json:"type_of_phone_id" struct:"type_of_phone_id" gorm:"foreignKey:type_of_phone_id"`
	TypeOfPhone        *TypeOfPhone         `json:"type_of_phone,omitempty" struct:"type_of_phone,omitempty" gorm:"foreignKey:type_of_phone_id"`
	Phone              *string              `json:"phone" struct:"phone"`
	FirstLogin         bool                 `json:"first_login" struct:"first_login"`
	IsEnabled          bool                 `json:"is_enabled" struct:"is_enabled"`
	IsSuspended        bool                 `json:"is_suspended" struct:"is_suspended"`
	CreatedAt          time.Time            `json:"created_at" struct:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at" struct:"updated_at"`
	DeletedAt          *gorm.DeletedAt      `json:"deleted_at" struct:"deleted_at"`
	AccountPermissions *[]AccountPermission `json:"account_permissions,omitempty" struct:"account_permissions,omitempty" gorm:"foreignKey:account_id"`
	AccountProperties  *[]AccountProperty   `json:"account_properties,omitempty" struct:"account_properties,omitempty" gorm:"foreignKey:account_id"`
	AuditLogAccounts   *[]AuditLogAccount   `json:"audit_log_accounts,omitempty" struct:"audit_log_accounts,omitempty" gorm:"foreignKey:account_id"`
	HistoryLogins      *[]HistoryLogin      `json:"history_logins,omitempty" struct:"history_logins,omitempty" gorm:"foreignKey:account_id"`
}
