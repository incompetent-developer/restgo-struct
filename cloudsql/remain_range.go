package cloudsql

import "time"

type RemainRange struct {
	ID        string    `json:"id" struct:"id" gorm:"primaryKey"`
	Name      string    `json:"name" struct:"name"`
	Min       int64     `json:"min" struct:"min"`
	Max       int64     `json:"max" struct:"max"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
