package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                         string          `json:"id" struct:"id" gorm:"primaryKey"`
	Firstname                  *string         `json:"firstname,omitempty" struct:"firstname,omitempty"`
	Lastname                   *string         `json:"lastname,omitempty" struct:"lastname,omitempty"`
	Account                    *string         `json:"account,omitempty" struct:"account,omitempty"`
	Password                   *string         `json:"-" struct:"password"`
	TypeOfGenderID             *string         `json:"type_of_gender_id,omitempty" struct:"type_of_gender_id,omitempty" gorm:"foreignKey:type_of_gender_id"`
	TypeOfGender               *TypeOfGender   `json:"type_of_gender,omitempty" struct:"type_of_gender,omitempty" gorm:"foreignKey:type_of_gender_id"`
	TypeOfUserID               string          `json:"type_of_user_id" struct:"type_of_user_id" gorm:"foreignKey:type_of_user_id"`
	TypeOfUser                 *TypeOfUser     `json:"type_of_user,omitempty" struct:"type_of_user,omitempty" gorm:"foreignKey:type_of_user_id"`
	TypeOfPhoneID              *string         `json:"type_of_phone_id,omitempty" struct:"type_of_phone_id,omitempty" gorm:"foreignKey:type_of_phone_id"`
	TypeOfPhone                *TypeOfPhone    `json:"type_of_phone,omitempty" struct:"type_of_phone,omitempty" gorm:"foreignKey:type_of_phone_id"`
	ContactPhone               *string         `json:"contact_phone,omitempty" struct:"contact_phone,omitempty"`
	ContactEmail               *string         `json:"contact_email,omitempty" struct:"contact_email,omitempty"`
	Ranking                    int64           `json:"ranking" struct:"ranking" gorm:"default:0"`
	Credit                     float64         `json:"credit" struct:"credit" gorm:"default:0"`
	FacebookID                 *string         `json:"-" struct:"facebook_id,omitempty"`
	FacebookAccount            *string         `json:"facebook_account,omitempty" struct:"facebook_account,omitempty"`
	GoogleID                   *string         `json:"-" struct:"google_id,omitempty"`
	GoogleAccount              *string         `json:"google_account,omitempty" struct:"google_account,omitempty"`
	AppleID                    *string         `json:"-" struct:"apple_id,omitempty"`
	AppleAccount               *string         `json:"apple_account,omitempty" struct:"apple_account,omitempty"`
	Birthdate                  *string         `json:"birthdate,omitempty" struct:"birthdate,omitempty"`
	TypeOfLanguageID           string          `json:"type_of_language_id" struct:"type_of_language_id" gorm:"foreignKey:type_of_language_id,default:thai"`
	TypeOfLanguage             *TypeOfLanguage `json:"type_of_language,omitempty" struct:"type_of_language,omitempty" gorm:"foreignKey:type_of_language_id"`
	TypeOfCurrencyID           string          `json:"type_of_currency_id" struct:"type_of_currency_id" gorm:"foreignKey:type_of_currency_id,default:thb"`
	TypeOfCurrency             *TypeOfCurrency `json:"type_of_currency,omitempty" struct:"type_of_currency,omitempty" gorm:"foreignKey:type_of_currency_id"`
	ImageUrl                   string          `json:"image_url" struct:"image_url"`
	ImageStorage               string          `json:"image_storage" struct:"image_storage"`
	Verify                     bool            `json:"verify" struct:"verify"`
	IsReceiveNotiCampaign      bool            `json:"is_receive_noti_campaign" struct:"is_receive_noti_campaign"`
	IsReceiveNotiBooking       bool            `json:"is_receive_noti_booking" struct:"is_receive_noti_booking"`
	IsReceiveNotiCredit        bool            `json:"is_receive_noti_credit" struct:"is_receive_noti_credit"`
	IsReceiveNotiReview        bool            `json:"is_receive_noti_review" struct:"is_receive_noti_review"`
	IsReceiveNotiEmailCampaign bool            `json:"is_receive_noti_email_campaign" struct:"is_receive_noti_email_campaign"`
	IsReceiveEmailNewsOffer    bool            `json:"is_receive_email_news_offer" struct:"is_receive_email_news_offer"`
	IsReceiveEmailPromotion    bool            `json:"is_receive_email_promotion" struct:"is_receive_email_promotion"`
	IsEnabled                  bool            `json:"is_enabled" struct:"is_enabled"`
	IsSuspended                bool            `json:"is_suspended" struct:"is_suspended"`
	IsRestgo                   bool            `json:"is_restgo" struct:"is_restgo"`
	RegisterDate               *time.Time      `json:"register_date" struct:"register_date"`
	CreatedAt                  time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt                  time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt                  gorm.DeletedAt  `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	UserFavorites              *[]UserFavorite `json:"user_favorites,omitempty" struct:"user_favorites,omitempty" gorm:"foreignKey:user_id"`
}
