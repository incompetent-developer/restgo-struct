package cloudsql

import "time"

// LineMessageLog structure for mysql
type LineMessageLog struct {
	ID                 int64      `json:"id" structs:"id" gorm:"primary_key;auto_increment:true"`
	LineConfigID       int64      `json:"line_config_id" structs:"line_config_id" `
	ReferenceType      string     `json:"reference_type" structs:"reference_type"`
	ReferenceID        string     `json:"reference_id" structs:"reference_id"`
	ReferenceCreatedAt *time.Time `json:"reference_created_at,omitempty" structs:"reference_created_at,omitempty"`
	LineNotiTypeID     string     `json:"line_noti_type_id" structs:"line_noti_type_id"`
	ErrorDescription   string     `json:"error_description" structs:"error_description"`
	CreatedAt          time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID           *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType         *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
