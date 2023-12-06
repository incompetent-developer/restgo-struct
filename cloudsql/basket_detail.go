package cloudsql

import "time"

type BasketDetail struct {
	BasketID  string    `json:"basket_id" struct:"basket_id" gorm:"foreignKey:basket_id"`
	Basket    *Basket   `json:"basket,omitempty" struct:"basket,omitempty" gorm:"foreignKey:basket_id"`
	RoomID    string    `json:"room_id" struct:"room_id" gorm:"foreignKey:room_id"`
	Room      *Room     `json:"room,omitempty" struct:"room,omitempty" gorm:"foreignKey:room_id"`
	CreatedAt time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt time.Time `json:"updated_at" struct:"updated_at"`
}
