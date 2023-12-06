package cloudsql

import "time"

// LineConfiguration structure for mysql
type LineConfiguration struct {
	ID                         int64                     `json:"id" structs:"id" gorm:"primary_key;auto_increment:true"`
	LinePropertyFollowerID     int64                     `json:"line_property_follower_id" structs:"line_property_follower_id" `
	LineNotificationTemplateID int64                     `json:"line_notification_template_id" structs:"line_notification_template_id"`
	LineNotificationTemplate   *LineNotificationTemplate `json:"line_notification_template,omitempty" structs:"line_notification_template,omitempty"`
	IsEnabled                  bool                      `json:"is_enabled" structs:"is_enabled"`
	CreatedAt                  time.Time                 `json:"created_at" structs:"created_at"`
	UpdatedAt                  time.Time                 `json:"updated_at" structs:"updated_at"`
	DeletedAt                  *time.Time                `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID                   *string                   `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType                 *string                   `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
