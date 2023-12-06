package cloudsql

import "time"

// LinePropertyFollower structure for mysql
type LinePropertyFollower struct {
	ID             int64         `json:"id" structs:"id" gorm:"primary_key;auto_increment:true"`
	LineFollowerID int64         `json:"line_follower_id" structs:"line_follower_id" `
	LineFollower   *LineFollower `json:"line_follower,omitempty" structs:"line_follower_id,omitempty" gorm:"PRELOAD:false"`
	PropertyID     string        `json:"property_id" structs:"property_id"`
	Property       *Property     `json:"property,omitempty" structs:"property,omitempty" gorm:"PRELOAD:false"`
	CreatedAt      time.Time     `json:"created_at" structs:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" structs:"updated_at"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID       *string       `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType     *string       `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
