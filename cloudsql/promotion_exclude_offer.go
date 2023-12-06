package cloudsql

import "time"

type PromotionExcludeOffer struct {
	PromotionID   string       `json:"promotion_id" struct:"promotion_id" gorm:"foreignKey:promotion_id"`
	Promotion     *Promotion   `json:"promotion,omitempty" struct:"promotion,omitempty" gorm:"foreignKey:promotion_id"`
	TypeOfOfferID string       `json:"type_of_offer_id" struct:"type_of_offer_id" gorm:"foreignKey:type_of_offer_id"`
	TypeOfOffer   *TypeOfOffer `json:"type_of_offer,omitempty" struct:"type_of_offer,omitempty" gorm:"foreignKey:type_of_offer_id"`
	CreatedAt     time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" struct:"updated_at"`
}
