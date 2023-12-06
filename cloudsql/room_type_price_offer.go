package cloudsql

import "time"

type RoomTypePriceOffer struct {
	RoomTypePriceID string         `json:"room_type_price_id" struct:"room_type_price_id" gorm:"foreignKey:room_type_price_id"`
	RoomTypePrice   *RoomTypePrice `json:"room_type_price,omitempty" struct:"room_type_price,omitempty" gorm:"foreignKey:room_type_price_id"`
	TypeOfOfferID   string         `json:"type_of_offer_id" struct:"type_of_offer_id" gorm:"foreignKey:type_of_offer_id"`
	TypeOfOffer     *TypeOfOffer   `json:"type_of_offer,omitempty" struct:"type_of_offer,omitempty" gorm:"foreignKey:type_of_offer_id"`
	StartDate       time.Time      `json:"start_date" struct:"start_date"`
	EndDate         *time.Time     `json:"end_date,omitempty" struct:"end_date,omitempty"`
	Forever         bool           `json:"forever" struct:"forever"`
	Price           float64        `json:"price" struct:"price"`
	Amount          int64          `json:"amount" struct:"amount"`
	CancelPolicyID  string         `json:"cancel_policy_id" struct:"cancel_policy_id" gorm:"foreignKey:cancel_policy_id"`
	CancelPolicy    *CancelPolicy  `json:"cancel_policy,omitempty" struct:"cancel_policy,omitempty" gorm:"foreignKey:cancel_policy_id"`
	IsEnabled       bool           `json:"is_enabled" struct:"is_enabled"`
	CreatedAt       time.Time      `json:"created_at" struct:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" struct:"updated_at"`
}
