package cloudsql

import "time"

// LineNotificationTemplate structure for mysql
type LineNotificationTemplate struct {
	ID              int64      `json:"id" structs:"id" gorm:"primary_key;auto_increment:true"`
	LineNotiTypeID  string     `json:"line_noti_type_id" structs:"line_noti_type_id"`
	TemplateMessage string     `json:"template_message" structs:"template_message"`
	LanguageID      string     `json:"language_id" struct:"language_id" `
	CreatedAt       time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID        *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType      *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
