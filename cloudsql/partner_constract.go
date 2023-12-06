package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type PartnerContract struct {
	ID                          string          `json:"id" struct:"id" gorm:"primaryKey"`
	FirstName                   string          `json:"first_name" struct:"first_name"`
	LastName                    string          `json:"last_name" struct:"last_name"`
	ContactEmail                string          `json:"contact_email" struct:"contact_email"`
	TypeOfContactPhoneID        *string         `json:"type_of_contact_phone_id,omitempty" struct:"type_of_contact_phone_id,omitempty" gorm:"foreignKey:type_of_contact_phone_id"`
	TypeOfContactPhone          *TypeOfPhone    `json:"type_of_contact_phone,omitempty" struct:"type_of_contact_phone,omitempty" gorm:"foreignKey:type_of_contact_phone_id;references:ID"`
	ContactPhone                string          `json:"contact_phone" struct:"contact_phone"`
	PropertyNameTH              string          `json:"property_name_th" struct:"property_name_th"`
	PropertyNameEN              *string         `json:"property_name_en,omitempty" struct:"property_name_en,omitempty"`
	TypeOfPropertyID            string          `json:"type_of_property_id" struct:"type_of_property_id" gorm:"foreignKey:type_of_property_id"`
	TypeOfProperty              *TypeOfProperty `json:"type_of_property,omitempty" struct:"type_of_property,omitempty" gorm:"foreignKey:type_of_property_id;references:ID"`
	NumberRoom                  int64           `json:"number_room" struct:"number_room"`
	Address                     string          `json:"address" struct:"address"`
	Province                    string          `json:"province" struct:"province"`
	District                    string          `json:"district" struct:"district"`
	SubDistrict                 string          `json:"sub_district" struct:"sub_district"`
	ZipCode                     string          `json:"zip_code" struct:"zip_code"`
	TypeOfPropertyOfficePhoneID *string         `json:"type_of_property_office_phone_id,omitempty" struct:"type_of_property_office_phone_id,omitempty" gorm:"foreignKey:type_of_property_office_phone_id"`
	TypeOfPropertyOfficePhone   *TypeOfPhone    `json:"type_of_property_office_phone,omitempty" struct:"type_of_property_office_phone,omitempty" gorm:"foreignKey:type_of_property_office_phone_id;references:ID"`
	PropertyOfficePhone         *string         `json:"property_office_phone,omitempty" struct:"property_office_phone,omitempty"`
	TypeOfPropertyMobilePhoneID *string         `json:"type_of_property_mobile_phone_id,omitempty" struct:"type_of_property_mobile_phone_id,omitempty" gorm:"foreignKey:type_of_property_mobile_phone_id"`
	TypeOfPropertyMobilePhone   *TypeOfPhone    `json:"type_of_property_mobile_phone,omitempty" struct:"type_of_property_mobile_phone,omitempty" gorm:"foreignKey:type_of_property_mobile_phone_id;references:ID"`
	PropertyMobilePhone         string          `json:"property_mobile_phone" struct:"property_mobile_phone"`
	Website                     *string         `json:"website,omitempty" struct:"website,omitempty"`
	SocialMedia                 *string         `json:"social_media,omitempty" struct:"social_media,omitempty"`
	IsReceiveEmailMarketing     bool            `json:"is_receive_email_marketing" struct:"is_receive_email_marketing"`
	IsEnabled                   bool            `json:"is_enabled" struct:"is_enabled"`
	IsSuspended                 bool            `json:"is_suspended" struct:"is_suspended"`
	CreatedAt                   time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt                   time.Time       `json:"updated_at" struct:"updated_at"`
	DeletedAt                   *gorm.DeletedAt `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
}
