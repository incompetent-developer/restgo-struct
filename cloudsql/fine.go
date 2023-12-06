package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Fine struct {
	ID                     string                `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID             string                `json:"property_id" struct:"property_id"`
	Property               *Property             `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	TypeOfMonthID          string                `json:"type_of_month_id" struct:"type_of_month_id" gorm:"foreignKey:type_of_month_id"`
	TypeOfMonth            *TypeOfMonth          `json:"type_of_month,omitempty" struct:"type_of_month,omitempty" gorm:"foreignKey:type_of_month_id"`
	Year                   int64                 `json:"year" struct:"year"`
	TotalFine              float64               `json:"total_fine" struct:"total_fine"`
	StatusFineOfTransferID string                `json:"status_fine_of_transfer_id" struct:"status_fine_of_transfer_id" gorm:"foreignKey:status_fine_of_transfer_id"` //ยังไม่ได้ชําระ, ชําระแล้ว, เกินกําหนดชําระ, รอการตรวจสอบ, ไม่ถูกต้อง
	StatusFineOfTransfer   *StatusFineOfTransfer `json:"status_fine_of_transfer,omitempty" struct:"status_fine_of_transfer,omitempty" gorm:"foreignKey:status_fine_of_transfer_id"`
	CreatedAt              time.Time             `json:"created_at" struct:"created_at"`
	UpdatedAt              time.Time             `json:"updated_at" struct:"updated_at"`
	DeletedAt              *gorm.DeletedAt       `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	FineBookings           *[]FineBooking        `json:"fine_bookings,omitempty" struct:"fine_bookings,omitempty" gorm:"foreignKey:fine_id"`
	FineImages             *[]FineImage          `json:"fine_images,omitempty" struct:"fine_images,omitempty" gorm:"foreignKey:fine_id"`
	FineLogs               *[]FineLog            `json:"fine_logs,omitempty" struct:"fine_logs,omitempty" gorm:"foreignKey:fine_id"`
}
