package cloudsql

import (
	"time"

	"gorm.io/gorm"
)

type Property struct {
	ID                    string              `json:"id" struct:"id" gorm:"primaryKey"`
	NameTH                string              `json:"name_th" struct:"name_th"`
	NameEN                *string             `json:"name_en" struct:"name_en"`
	TypeOfPropertyID      string              `json:"type_of_property_id" struct:"type_of_property_id" gorm:"foreignKey:type_of_property_id"`
	TypeOfProperty        *TypeOfProperty     `json:"type_of_property,omitempty" struct:"type_of_property,omitempty" gorm:"foreignKey:type_of_property_id"`
	Level                 int64               `json:"level" struct:"level" gorm:"default:0"`
	TypeOfPhoneID         *string             `json:"type_of_phone_id" struct:"type_of_phone_id" gorm:"foreignKey:type_of_phone_id"`
	TypeOfPhone           *TypeOfPhone        `json:"type_of_phone,omitempty" struct:"type_of_phone,omitempty" gorm:"foreignKey:type_of_phone_id"`
	Phone                 *string             `json:"phone" struct:"phone"`
	Fax                   *string             `json:"fax" struct:"fax"`
	Email                 *string             `json:"email" struct:"email"`
	Email2                *string             `json:"email2" struct:"email2"`
	EmailInformation      *string             `json:"email_information" struct:"email_information"`
	PreDayBook            int64               `json:"pre_day_book" struct:"pre_day_book"`
	Website               *string             `json:"website" struct:"website"`
	Facebook              *string             `json:"facebook" struct:"facebook"`
	ChildAgeRange         int64               `json:"child_age_range" struct:"child_age_range" gorm:"default:1"`
	AdultAgeRange         int64               `json:"adult_age_range" struct:"adult_age_range" gorm:"default:1"`
	Reception             bool                `json:"reception" struct:"reception"`
	ReceptionOpen         *string             `json:"reception_open" struct:"reception_open"`
	ReceptionClose        *string             `json:"reception_close" struct:"reception_close"`
	CheckinTime           *string             `json:"checkin_time" struct:"checkin_time"`
	CheckoutTime          *string             `json:"checkout_time" struct:"checkout_time"`
	TypeOfCurrencyID      *string             `json:"type_of_currency_id" struct:"type_of_currency_id" gorm:"foreignKey:type_of_currency_id; default:thb"`
	TypeOfCurrency        *TypeOfCurrency     `json:"type_of_currency,omitempty" struct:"type_of_currency,omitempty" gorm:"foreignKey:type_of_currency_id"`
	Breakfast             bool                `json:"breakfast" struct:"breakfast"`
	BreakfastOpen         *string             `json:"breakfast_open" struct:"breakfast_open"`
	BreakfastClose        *string             `json:"breakfast_close" struct:"breakfast_close"`
	Pool                  bool                `json:"pool" struct:"pool"`
	PoolOpen              *string             `json:"pool_open" struct:"pool_open"`
	PoolClose             *string             `json:"pool_close" struct:"pool_close"`
	Pet                   bool                `json:"pet" struct:"pet"`
	IsPetCost             *bool               `json:"is_pet_cost" struct:"is_pet_cost"`
	IsPetSpecifySize      *bool               `json:"is_pet_specify_size" struct:"is_pet_specify_size"`
	PetTypeOfPriceID      *string             `json:"pet_type_of_price_id" struct:"pet_type_of_price_id" gorm:"foreignKey:pet_type_of_price_id"`
	PetTypeOfPrice        *PetTypeOfPrice     `json:"pet_type_of_price" struct:"pet_type_of_price" gorm:"foreignKey:pet_type_of_price_id"`
	PetDescription        *string             `json:"pet_description" struct:"pet_description"`
	OtherInformation      *string             `json:"other_information" struct:"other_information"`
	Announce              *string             `json:"announce" struct:"announce"`
	Description           *string             `json:"description" struct:"description"`
	BlueprintImageUrl     *string             `json:"blueprint_image_url" struct:"blueprint_image_url"`
	BlueprintImageStorage *string             `json:"blueprint_image_storage" struct:"blueprint_image_storage"`
	AccountID             string              `json:"account_id" struct:"account_id" gorm:"foreignKey:account_id"`
	Account               *Account            `json:"account,omitempty" struct:"account,omitempty" gorm:"foreignKey:account_id"`
	CancelPolicyID        *string             `json:"cancel_policy_id" struct:"cancel_policy_id" gorm:"foreignKey:cancel_policy_id"`
	CancelPolicy          *CancelPolicy       `json:"cancel_policy,omitempty" struct:"cancel_policy,omitempty" gorm:"foreignKey:cancel_policy_id"`
	TaxID                 string              `json:"tax_id" struct:"tax_id"`                                          // หมายเลขประจำตัวผู้เสียภาษี
	TaxAccountName        *string             `json:"tax_account_name,omitempty" struct:"tax_account_name,omitempty" ` // ชื่อผู้เสียภาษี
	IsBookingNow          bool                `json:"is_booking_now" struct:"is_booking_now"`                          // จองทันที = true, ปกติ = false
	IsPrepay              bool                `json:"is_prepay" struct:"is_prepay"`                                    // ชำระล่วงหน้า = true, ปกติ = false
	IsPayOnline           bool                `json:"is_pay_online" struct:"is_pay_online"`                            // default true
	IsPayAt               bool                `json:"is_pay_at" struct:"is_pay_at"`                                    // default false
	IsRecommended         bool                `json:"is_recommended" struct:"is_recommended"`                          // แนะนำรึยัง = true, ยัง = false
	IsCorporation         bool                `json:"is_corporation" struct:"is_corporation"`                          // เป็นนิติบุคคล = true, บุคคลธรรมดา = false
	IsApplyVat            bool                `json:"is_apply_vat" struct:"is_apply_vat"`                              // เข้าร่วมภาษีมูลค่าเพิ่ม = true, ไม่เข้าร่วม = false
	IsEnabled             bool                `json:"is_enabled" struct:"is_enabled"`                                  // publish ในหน้าแรก หรือไม้
	IsSuspended           bool                `json:"is_suspended" struct:"is_suspended"`                              // active อยู่หรือไม่
	IsBookingEnabled      bool                `json:"is_booking_enabled" struct:"is_booking_enabled"`                  // เปิดให้ จอง ที่พักได้หรือไม่?
	IsRestgo              bool                `json:"is_restgo" struct:"is_restgo"`
	CreatedAt             time.Time           `json:"created_at" struct:"created_at"`
	UpdatedAt             time.Time           `json:"updated_at" struct:"updated_at"`
	DeletedAt             *gorm.DeletedAt     `json:"deleted_at,omitempty" struct:"deleted_at,omitempty"`
	PropertyPets          *[]PropertyPet      `json:"property_pets,omitempty" struct:"property_pets,omitempty" gorm:"foreignKey:property_id"`
	PropertyAddress       *PropertyAddress    `json:"property_address,omitempty" struct:"property_address,omitempty" gorm:"foreignKey:property_id"`
	PropertyContacts      *[]PropertyContact  `json:"property_contacts,omitempty" struct:"property_contacts,omitempty" gorm:"foreignKey:property_id"`
	PropertyBanks         *[]PropertyBank     `json:"property_banks,omitempty" struct:"property_banks,omitempty" gorm:"foreignKey:property_id"`
	PropertyImages        *[]PropertyImage    `json:"property_images,omitempty" struct:"property_images,omitempty" gorm:"foreignKey:property_id"`
	PropertyFacilities    *[]PropertyFacility `json:"property_facilities,omitempty" struct:"property_facilities,omitempty" gorm:"foreignKey:property_id"`
	RoomTypes             *[]RoomType         `json:"room_types,omitempty" struct:"room_types,omitempty" gorm:"foreignKey:property_id"`
	Reviews               *[]Review           `json:"reviews,omitempty" struct:"reviews,omitempty" gorm:"foreignKey:property_id"`
}
