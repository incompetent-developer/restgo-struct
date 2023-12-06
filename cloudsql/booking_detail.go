package cloudsql

import "time"

type BookingDetail struct {
	BookingID         string          `json:"booking_id" struct:"booking_id" gorm:"foreignKey:booking_id"`
	Booking           *Booking        `json:"booking,omitempty" struct:"booking,omitempty" gorm:"foreignKey:booking_id"`
	RoomID            string          `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room              *Room           `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	RoomBeds          string          `json:"room_beds" struct:"room_beds"`
	TypeOfOfferID     *string         `json:"type_of_offer_id,omitempty" struct:"type_of_offer_id,omitempty" gorm:"foreignKey:type_of_offer_id"`
	TypeOfOffer       *TypeOfOffer    `json:"type_of_offer,omitempty" struct:"type_of_offer,omitempty" gorm:"foreignKey:type_of_offer_id"`
	PromotionID       *string         `json:"promotion_id,omitempty" struct:"promotion_id,omitempty" gorm:"foreignKey:promotion_id"`
	Promotion         *Promotion      `json:"promotion,omitempty" struct:"promotion,omitempty" gorm:"foreignKey:promotion_id"`
	PromotionName     *string         `json:"promotion_name,omitempty" struct:"promotion_name,omitempty"`
	PromotionDiscount *float64        `json:"promotion_discount,omitempty" struct:"promotion_discount,omitempty" gorm:"default:0"`
	TypeOfDiscountID  *string         `json:"type_of_discount_id,omitempty" struct:"type_of_discount_id,omitempty"`
	TypeOfDiscount    *TypeOfDiscount `json:"type_of_discount,omitempty" struct:"type_of_discount,omitempty" gorm:"foreignKey:type_of_discount_id"`
	BasePrice         float64         `json:"base_price" struct:"base_price" gorm:"default:0"`
	AddonPrice        float64         `json:"addon_price" struct:"addon_price" gorm:"default:0"`
	AddonPercentage   float64         `json:"addon_percentage" struct:"addon_percentage" gorm:"default:0"`
	NumberExtraBed    int64           `json:"number_extra_bed" struct:"number_extra_bed" gorm:"default:0"`
	ExtraBedPrice     float64         `json:"extra_bed_price" struct:"extra_bed_price" gorm:"default:0"`
	OfferPrice        float64         `json:"offer_price" struct:"offer_price" gorm:"default:0"`
	DiscountPrice     float64         `json:"discount_price" struct:"discount_price" gorm:"default:0"`
	FinalPrice        float64         `json:"final_price" struct:"final_price" gorm:"default:0"`
	CreatedAt         time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" struct:"updated_at"`
}
