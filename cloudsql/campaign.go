package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	ID                           string                         `json:"id" struct:"id" gorm:"primaryKey"`
	Name                         string                         `json:"name" struct:"name"`
	ImageUrl                     *string                        `json:"image_url,omitempty" struct:"image_url,omitempty"`
	ImageStorage                 *string                        `json:"image_storage,omitempty" struct:"image_storage,omitempty"`
	CampaignStartDate            *time.Time                     `json:"campaign_start_date,omitempty" struct:"campaign_start_date,omitempty"`
	CampaignEndDate              *time.Time                     `json:"campaign_end_date,omitempty" struct:"campaign_end_date,omitempty"`
	CampaignForever              bool                           `json:"campaign_forever" struct:"campaign_forever"`
	StayStartDate                *time.Time                     `json:"stay_start_date,omitempty" struct:"stay_start_date,omitempty"`
	StayEndDate                  *time.Time                     `json:"stay_end_date,omitempty" struct:"stay_end_date,omitempty"`
	StayForever                  bool                           `json:"stay_forever" struct:"stay_forever"`
	RoomMin                      *int64                         `json:"room_min,omitempty" struct:"room_min,omitempty"`
	RoomMax                      *int64                         `json:"room_max,omitempty" struct:"room_max,omitempty"`
	StayMin                      *int64                         `json:"stay_min,omitempty" struct:"stay_min,omitempty"`
	StayMax                      *int64                         `json:"stay_max,omitempty" struct:"stay_max,omitempty"`
	BudgetPrice                  float64                        `json:"budget_price" struct:"budget_price"`
	BudgetPriceDayLimit          *float64                       `json:"budget_price_day_limit,omitempty" struct:"budget_price_day_limit,omitempty"`
	TypeOfUserJoinID             string                         `json:"type_of_user_join_id" struct:"type_of_user_join_id" gorm:"foreignKey:type_of_user_join_id"`
	TypeOfUserJoin               *TypeOfUserJoin                `json:"type_of_user_join,omitempty" struct:"type_of_user_join,omitempty" gorm:"foreignKey:type_of_user_join_id"`
	UserAgeMin                   *int64                         `json:"user_age_min,omitempty" struct:"user_age_min,omitempty"`
	UserAgeMax                   *int64                         `json:"user_age_max,omitempty" struct:"user_age_max,omitempty"`
	UserCreditTypeOfCompareID    *string                        `json:"user_credit_type_of_compare_id" struct:"user_credit_type_of_compare_id" gorm:"foreignKey:user_credit_type_of_compare_id"`
	UserCreditTypeOfCompare      *TypeOfCompare                 `json:"user_credit_type_of_compare,omitempty" struct:"user_credit_type_of_compare,omitempty" gorm:"foreignKey:user_credit_type_of_compare_id"`
	UserCreditMin                *float64                       `json:"user_credit_min,omitempty" struct:"user_credit_min,omitempty"`
	UserCreditMax                *float64                       `json:"user_credit_max,omitempty" struct:"user_credit_max,omitempty"`
	UserBookingTypeOfCompareID   *string                        `json:"user_booking_type_of_compare_id" struct:"user_booking_type_of_compare_id" gorm:"foreignKey:user_booking_type_of_compare_id"`
	UserBookingTypeOfCompare     *TypeOfCompare                 `json:"user_booking_type_of_compare,omitempty" struct:"user_booking_type_of_compare,omitempty" gorm:"foreignKey:user_booking_type_of_compare_id"`
	UserBookingMin               *int64                         `json:"user_booking_min,omitempty" struct:"user_booking_min,omitempty"`
	UserBookingMax               *int64                         `json:"user_booking_max,omitempty" struct:"user_booking_max,omitempty"`
	UserCreatedStart             *time.Time                     `json:"user_created_start,omitempty" struct:"user_created_start,omitempty"`
	UserCreatedEnd               *time.Time                     `json:"user_created_end,omitempty" struct:"user_created_end,omitempty"`
	TypeOfPropertyJoinID         string                         `json:"type_of_property_join_id" struct:"type_of_property_join_id" gorm:"foreignKey:type_of_property_join_id"`
	TypeOfPropertyJoin           *TypeOfPropertyJoin            `json:"type_of_property_join,omitempty" struct:"type_of_property_join,omitempty" gorm:"foreignKey:type_of_property_join_id"`
	TypeOfDiscountID             string                         `json:"type_of_discount_id" struct:"type_of_discount_id" gorm:"foreignKey:type_of_discount_id"`
	TypeOfDiscount               *TypeOfDiscount                `json:"type_of_discount,omitempty" struct:"type_of_discount,omitempty" gorm:"foreignKey:type_of_discount_id"`
	Discount                     *float64                       `json:"discount,omitempty" struct:"discount,omitempty"`
	DiscountLimit                *float64                       `json:"discount_limit,omitempty" struct:"discount_limit,omitempty"`
	CreditExchange               *float64                       `json:"credit_exchange,omitempty" struct:"credit_exchange,omitempty"`
	NumberSale                   *int64                         `json:"number_sale,omitempty" struct:"number_sale,omitempty"`
	NumberRight                  *int64                         `json:"number_right,omitempty" struct:"number_right,omitempty"`
	CancelPolicyID               string                         `json:"cancel_policy_id" struct:"cancel_policy_id" gorm:"foreignKey:cancel_policy_id"`
	CancelPolicy                 *CancelPolicy                  `json:"cancel_policy,omitempty" struct:"cancel_policy,omitempty" gorm:"foreignKey:cancel_policy_id"`
	IsParticipatePromotion       bool                           `json:"is_participate_promotion" struct:"is_participate_promotion"`
	IsParticipateCampaign        bool                           `json:"is_participate_campaign" struct:"is_participate_campaign"`
	IsEnabled                    bool                           `json:"is_enabled" struct:"is_enabled"`
	IsSuspended                  bool                           `json:"is_suspended" struct:"is_suspended"`
	UpdatedTime                  time.Time                      `json:"updated_time" struct:"updated_time"`
	CreatedAt                    time.Time                      `json:"created_at" struct:"created_at"`
	UpdatedAt                    time.Time                      `json:"updated_at" struct:"updated_at"`
	DeletedAt                    *gorm.DeletedAt                `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	CampaignTimeSales            *[]CampaignTimeSale            `json:"campaign_time_sales,omitempty" struct:"campaign_time_sales,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignBookDays             *[]CampaignBookDay             `json:"campaign_book_days,omitempty" struct:"campaign_book_days,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignStayDays             *[]CampaignStayDay             `json:"campaign_stay_days,omitempty" struct:"campaign_stay_days,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignBirthdates           *[]CampaignBirthdate           `json:"campaign_birthdates,omitempty" struct:"campaign_birthdates,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignUserSpecifies        *[]CampaignUserSpecify         `json:"campaign_user_specifies,omitempty" struct:"campaign_user_specifies,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignMonths               *[]CampaignMonth               `json:"campaign_months,omitempty" struct:"campaign_months,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignTypeOfProperties     *[]CampaignTypeOfProperty      `json:"campaign_type_of_properties,omitempty" struct:"campaign_type_of_properties,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPropertyRegions      *[]CampaignPropertyRegion      `json:"campaign_property_regions,omitempty" struct:"campaign_property_regions,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPropertyProvinces    *[]CampaignPropertyProvince    `json:"campaign_property_provinces,omitempty" struct:"campaign_property_provinces,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPropertyDistricts    *[]CampaignPropertyDistrict    `json:"campaign_property_districts,omitempty" struct:"campaign_property_districts,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPropertySubDistricts *[]CampaignPropertySubDistrict `json:"campaign_property_sub_districts,omitempty" struct:"campaign_property_sub_districts,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignPropertySpecifies    *[]CampaignPropertySpecify     `json:"campaign_property_specifies,omitempty" struct:"campaign_property_specifies,omitempty" gorm:"foreignKey:campaign_id"`
	CampaignCodes                *[]CampaignCode                `json:"campaign_codes,omitempty" struct:"campaign_codes,omitempty" gorm:"foreignKey:campaign_id"`
}
