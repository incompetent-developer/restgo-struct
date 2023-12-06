package cloudsql

import "time"

type KsherFeeRate struct {
	ID            string     `json:"id" structs:"id"`
	FeePercentage *float64   `json:"fee_percentage,omitempty" structs:"fee_percentage,omitempty"`
	FeeValue      *float64   `json:"fee_value,omitempty" structs:"fee_value,omitempty"`
	CreatedAt     time.Time  `json:"created_at" structs:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" structs:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID      *string    `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType    *string    `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
