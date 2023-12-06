package cloudsql

import "time"

type TempBooking struct {
	ID                   string                 `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID           string                 `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property             *Property              `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	UserID               string                 `json:"user_id" struct:"user_id" gorm:"foreignKey:user_id"`
	User                 *User                  `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	NumberAdult          int64                  `json:"number_adult" struct:"number_adult" gorm:"default:1"`
	NumberChild          int64                  `json:"number_child" struct:"number_child" gorm:"default:0"`
	ContactFirstname     *string                `json:"contact_firstname,omitempty" struct:"contact_firstname,omitempty"`
	ContactLastname      *string                `json:"contact_lastname,omitempty" struct:"contact_lastname,omitempty"`
	ContactTypeOfPhoneID *string                `json:"contact_type_of_phone_id" struct:"contact_type_of_phone_id" gorm:"foreignKey:contact_type_of_phone_id"`
	ContactTypeOfPhone   *TypeOfPhone           `json:"contact_type_of_phone,omitempty" struct:"contact_type_of_phone,omitempty" gorm:"foreignKey:contact_type_of_phone_id"`
	ContactPhone         *string                `json:"contact_phone,omitempty" struct:"contact_phone,omitempty"`
	ContactEmail         *string                `json:"contact_email,omitempty" struct:"contact_email,omitempty"`
	BookingOther         bool                   `json:"booking_other" struct:"booking_other"`
	GuestFirstname       *string                `json:"guest_firstname,omitempty" struct:"guest_firstname,omitempty"`
	GuestLastname        *string                `json:"guest_lastname,omitempty" struct:"guest_lastname,omitempty"`
	GuestTypeOfPhoneID   *string                `json:"guest_type_of_phone_id" struct:"guest_type_of_phone_id" gorm:"foreignKey:guest_type_of_phone_id"`
	GuestTypeOfPhone     *TypeOfPhone           `json:"guest_type_of_phone,omitempty" struct:"guest_type_of_phone,omitempty" gorm:"foreignKey:guest_type_of_phone_id"`
	GuestPhone           *string                `json:"guest_phone,omitempty" struct:"guest_phone,omitempty"`
	CheckinDate          time.Time              `json:"checkin_date" struct:"checkin_date"`
	CheckoutDate         time.Time              `json:"checkout_date" struct:"checkout_date"`
	NumberNight          int64                  `json:"number_night" struct:"number_night"`
	RoomRequest          *string                `json:"room_request,omitempty" struct:"room_request,omitempty"`
	BedRequest           *string                `json:"bed_request,omitempty" struct:"bed_request,omitempty"`
	TimeRequest          *string                `json:"time_request,omitempty" struct:"time_request,omitempty"`
	OtherRequest         *string                `json:"other_request,omitempty" struct:"other_request,omitempty"`
	RoomRate             float64                `json:"room_rate" struct:"room_rate" gorm:"default:0"`
	ChangeFee            float64                `json:"change_fee" struct:"change_fee" gorm:"default:0"`
	Discount             float64                `json:"discount" struct:"discount" gorm:"default:0"`
	CampaignPrice        float64                `json:"campaign_price" struct:"campaign_price"`
	Vat                  float64                `json:"vat" struct:"vat" gorm:"default:0"`
	TotalBasePrice       float64                `json:"total_base_price" struct:"total_base_price" gorm:"default:0"`
	TotalAddonPrice      float64                `json:"total_addon_price" struct:"total_addon_price" gorm:"default:0"`
	TotalExtraBedPrice   float64                `json:"total_extra_bed_price" struct:"total_extra_bed_price"`
	TotalPetPrice        float64                `json:"total_pet_price" struct:"total_pet_price"`
	NetPrice             float64                `json:"net_price" struct:"net_price" gorm:"default:0"`
	CreditExchange       float64                `json:"credit_exchange" struct:"credit_exchange" gorm:"default:0"`
	StatusOfBookingID    string                 `json:"status_of_booking_id" struct:"status_of_booking_id" gorm:"foreignKey:status_of_booking_id"`
	StatusOfBooking      *StatusOfBooking       `json:"status_of_booking,omitempty" struct:"status_of_booking,omitempty" gorm:"foreignKey:status_of_booking_id"`
	CancelPolicyID       string                 `json:"cancel_policy_id" struct:"cancel_policy_id" gorm:"foreignKey:cancel_policy_id"`
	CancelPolicy         *CancelPolicy          `json:"cancel_policy,omitempty" struct:"cancel_policy,omitempty" gorm:"foreignKey:cancel_policy_id"`
	CancelPolicyDate     *time.Time             `json:"cancel_policy_date,omitempty" struct:"cancel_policy_date,omitempty"`
	CancelDate           *time.Time             `json:"cancel_date,omitempty" struct:"cancel_date,omitempty"`
	ReasonOfCancelID     *string                `json:"reason_of_cancel_id,omitempty" struct:"reason_of_cancel_id,omitempty"`
	ReasonOfCancel       *ReasonOfCancel        `json:"reason_of_cancel,omitempty" struct:"reason_of_cancel,omitempty" gorm:"foreignKey:reason_of_cancel_id"`
	CancelDescription    *string                `json:"cancel_description,omitempty" struct:"cancel_description,omitempty"`
	CancelAmountPrice    *float64               `json:"cancel_amount_price,omitempty" struct:"cancel_amount_price,omitempty"`
	CancelRefundCredit   *float64               `json:"cancel_refund_credit,omitempty" struct:"cancel_refund_credit,omitempty"`
	CancelerID           *string                `json:"canceler_id,omitempty" struct:"canceler_id,omitempty"`
	IsRead               bool                   `json:"is_read" struct:"is_read"`
	IsPet                bool                   `json:"is_pet" struct:"is_pet"`
	IsReview             bool                   `json:"is_review" struct:"is_review"`
	IsUpdateBooking      bool                   `json:"is_update_booking" struct:"is_update_booking"`
	IsEnabled            bool                   `json:"is_enabled" struct:"is_enabled"`
	IsSuspended          bool                   `json:"is_suspended" struct:"is_suspended"`
	ReferenceBookingID   *string                `json:"reference_booking_id,omitempty" struct:"reference_booking_id,omitempty"`
	BookingPaymentTypeID *string                `json:"booking_payment_type_id,omitempty" struct:"booking_payment_type_id,omitempty" gorm:"foreignKey:booking_payment_type_id"`
	BookingPaymentType   *BookingPaymentType    `json:"booking_payment_type,omitempty" struct:"booking_payment_type,omitempty" gorm:"foreignKey:booking_payment_type_id"`
	CreatedAt            time.Time              `json:"created_at" struct:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at" struct:"updated_at"`
	TempBookingDetails   *[]TempBookingDetail   `json:"temp_booking_details,omitempty" struct:"temp_booking_details,omitempty" gorm:"foreignKey:temp_booking_id"`
	TempBookingBaskets   *[]TempBookingBasket   `json:"temp_booking_baskets,omitempty" struct:"temp_booking_baskets,omitempty" gorm:"foreignKey:temp_booking_id"`
	TempBookingCampaigns *[]TempBookingCampaign `json:"temp_booking_campaigns,omitempty" struct:"temp_booking_campaigns,omitempty" gorm:"foreignKey:temp_booking_id"`
	TempBookingPets      *[]TempBookingPet      `json:"temp_booking_pets,omitempty" struct:"temp_booking_pets,omitempty" gorm:"foreignKey:temp_booking_id"`
	TempBookingPayment   *TempBookingPayment    `json:"temp_booking_payment,omitempty" struct:"temp_booking_payment,omitempty" gorm:"foreignKey:temp_booking_id"`
}
