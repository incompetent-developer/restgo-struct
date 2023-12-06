package cloudsql

import "time"

type UserNotification struct {
	ID                   string              `json:"id" struct:"id" gorm:"primaryKey"`
	UserID               string              `json:"user_id" struct:"user_id" gorm:"foreignKey:user_id"`
	User                 *User               `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	TypeOfNotificationID string              `json:"type_of_notification_id" struct:"type_of_notification_id" gorm:"foreignKey:type_of_notification_id"`
	TypeOfNotification   *TypeOfNotification `json:"type_of_notification,omitempty" struct:"type_of_notification,omitempty" gorm:"foreignKey:type_of_notification_id"`
	BookingID            *string             `json:"booking_id" struct:"booking_id" `
	PropertyID           *string             `json:"property_id" struct:"property_id"`
	PropertyNameTH       *string             `json:"property_name_th" struct:"property_name_th"`
	PropertyNameEN       *string             `json:"property_name_en " struct:"property_name_en"`
	CampaignID           *string             `json:"campaign_id,omitempty" struct:"campaign_id,omitempty" `
	CampaignName         *string             `json:"campaign_name" struct:"campaign_name"`
	CampaignType         *string             `json:"campaign_type" struct:"campaign_type"`
	CampaignDiscount     *float64            `json:"campaign_discount" struct:"campaign_discount"`
	StatusOfNotification *string             `json:"status_of_notification" struct:"status_of_notification"`
	Credit               *float64            `json:"credit" struct:"credit"`
	CreditExpireDate     *time.Time          `json:"credit_expire_date" struct:"credit_expire_date"`
	IsRead               bool                `json:"is_read" struct:"is_read"`
	CreatedAt            time.Time           `json:"created_at" struct:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at" struct:"updated_at"`
}
