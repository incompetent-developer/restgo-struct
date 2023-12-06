package cloudsql

import "time"

type Review struct {
	ID                        string          `json:"id" struct:"id" gorm:"primaryKey"`
	BookingID                 string          `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking                   *Booking        `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	UserID                    string          `json:"user_id" struct:"user_id" gorm:"foreignKey:user_id"`
	User                      *User           `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	PropertyID                string          `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property                  *Property       `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	TypeOfReviewID            *string         `json:"type_of_review_id" struct:"type_of_review_id" gorm:"foreignKey:type_of_review_id"`
	TypeOfReview              *TypeOfReview   `json:"type_of_review,omitempty" struct:"type_of_review,omitempty" gorm:"foreignKey:type_of_review_id"`
	Title                     *string         `json:"title,omitempty" struct:"title,omitempty"`
	Description               *string         `json:"description,omitempty" struct:"description,omitempty"`
	PointClean                float64         `json:"point_clean" struct:"point_clean" gorm:"default:0"`
	PointService              float64         `json:"point_service" struct:"point_service" gorm:"default:0"`
	PointWorth                float64         `json:"point_worth" struct:"point_worth" gorm:"default:0"`
	PointConvenient           float64         `json:"point_convenient" struct:"point_convenient" gorm:"default:0"`
	PointLocation             float64         `json:"point_location" struct:"point_location" gorm:"default:0"`
	StatusOfReviewID          *string         `json:"status_of_review_id" struct:"status_of_review_id" gorm:"foreignKey:status_of_review_id"`
	StatusOfReview            *StatusOfReview `json:"status_of_review" struct:"status_of_review" gorm:"foreignKey:status_of_review_id"`
	ReasonOfReportID          *string         `json:"reason_of_report_id" struct:"reason_of_report_id" gorm:"foreignKey:reason_of_report_id"`
	ReasonOfReport            *ReasonOfReport `json:"reason_of_report" struct:"reason_of_report" gorm:"foreignKey:reason_of_report_id"`
	ReasonOfReportDescription *string         `json:"reason_of_report_description" struct:"reason_of_report_description"`
	IsEnabled                 bool            `json:"is_enabled" struct:"is_enabled"`
	CreatedAt                 time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt                 time.Time       `json:"updated_at" struct:"updated_at"`
	ReviewImages              *[]ReviewImage  `json:"review_images,omitempty" struct:"review_images,omitempty" gorm:"foreignKey:review_id"`
}
