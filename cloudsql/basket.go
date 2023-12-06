package cloudsql

import "time"

type Basket struct {
	ID            string          `json:"id" struct:"id"`
	UserID        string          `json:"user_id" struct:"user_id" gorm:"foreignKey:user_id"`
	User          *User           `json:"user,omitempty" struct:"user,omitempty" gorm:"foreignKey:user_id"`
	PromotionID   *string         `json:"promotion_id,omitempty" struct:"promotion_id,omitempty" gorm:"foreignKey:promotion_id"`
	Promotion     *Promotion      `json:"promotion,omitempty" struct:"promotion,omitempty" gorm:"foreignKey:promotion_id"`
	TypeOfOfferID *string         `json:"type_of_offer_id,omitempty" struct:"type_of_offer_id,omitempty" gorm:"foreignKey:type_of_offer_id"`
	TypeOfOffer   *TypeOfOffer    `json:"type_of_offer,omitempty" struct:"type_of_offer,omitempty" gorm:"foreignKey:type_of_offer_id"`
	CheckIn       time.Time       `json:"check_in" struct:"check_in"`
	CheckOut      time.Time       `json:"check_out" struct:"check_out"`
	NumberAdult   int64           `json:"number_adult" struct:"number_adult" gorm:"default:1"`
	NumberChild   int64           `json:"number_child" struct:"number_child" gorm:"default:0"`
	CreatedAt     time.Time       `json:"created_at" struct:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" struct:"updated_at"`
	BasketDetails *[]BasketDetail `json:"basket_details,omitempty" struct:"basket_details,omitempty" gorm:"foreignKey:basket_id"`
}
