package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	ID                     string                   `json:"id" struct:"id" gorm:"primaryKey"`
	Name                   string                   `json:"name" struct:"name"`
	PropertyID             string                   `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property               *Property                `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	PromotionStartDate     *time.Time               `json:"promotion_start_date" struct:"promotion_start_date"`
	PromotionEndDate       *time.Time               `json:"promotion_end_date" struct:"promotion_end_date"`
	PromotionForever       bool                     `json:"promotion_forever" struct:"promotion_forever"`
	StayStartDate          *time.Time               `json:"stay_start_date" struct:"stay_start_date"`
	StayEndDate            *time.Time               `json:"stay_end_date" struct:"stay_end_date"`
	StayForever            bool                     `json:"stay_forever" struct:"stay_forever"`
	StayMin                *int64                   `json:"stay_min" struct:"stay_min"`
	StayMax                *int64                   `json:"stay_max" struct:"stay_max"`
	TypeOfDiscountID       string                   `json:"type_of_discount_id" struct:"type_of_discount_id" gorm:"foreignKey:type_of_discount_id"`
	TypeOfDiscount         *TypeOfDiscount          `json:"type_of_discount" struct:"type_of_discount" gorm:"foreignKey:type_of_discount_id"`
	Discount               *float64                 `json:"discount" struct:"discount"`
	CancelPolicyID         string                   `json:"cancel_policy_id" struct:"cancel_policy_id" gorm:"foreignKey:cancel_policy_id"`
	CancelPolicy           *CancelPolicy            `json:"cancel_policy" struct:"cancel_policy" gorm:"foreignKey:cancel_policy_id"`
	NumberSale             *int64                   `json:"number_sale" struct:"number_sale"`
	NumberRight            *int64                   `json:"number_right" struct:"number_right"`
	IsParticipate          bool                     `json:"is_participate" struct:"is_participate"`
	IsEnabled              bool                     `json:"is_enabled" struct:"is_enabled"`
	IsSuspended            bool                     `json:"is_suspended" struct:"is_suspended"`
	CreatedAt              time.Time                `json:"created_at" struct:"created_at"`
	UpdatedAt              time.Time                `json:"updated_at" struct:"updated_at"`
	DeletedAt              *gorm.DeletedAt          `json:"deleted_at" struct:"deleted_at"`
	PromotionRoomTypes     *[]PromotionRoomType     `json:"promotion_room_types,omitempty" struct:"promotion_room_types,omitempty" gorm:"foreignKey:promotion_id"`
	PromotionBookDays      *[]PromotionBookDay      `json:"promotion_book_days,omitempty" struct:"promotion_book_days,omitempty" gorm:"foreignKey:promotion_id"`
	PromotionStayDays      *[]PromotionStayDay      `json:"promotion_stay_days,omitempty" struct:"promotion_stay_days,omitempty" gorm:"foreignKey:promotion_id"`
	PromotionExcludeOffers *[]PromotionExcludeOffer `json:"promotion_exclude_offers,omitempty" struct:"promotion_exclude_offers,omitempty" gorm:"foreignKey:promotion_id"`
}
