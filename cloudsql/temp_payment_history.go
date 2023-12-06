package cloudsql

import "time"

type TempPaymentHistory struct {
	ID             string    `json:"id" struct:"id" gorm:"primaryKey"`
	ChannelOrderID *string   `json:"channel_order_id,omitempty" struct:"channel_order_id,omitempty"`
	OrderType      string    `json:"order_type" struct:"order_type"`
	OrderDate      *string   `json:"order_date,omitempty" struct:"order_date,omitempty"`
	ErrorMessage   string    `json:"error_message" struct:"error_message" `
	Timestamp      string    `json:"timestamp" struct:"timestamp"`
	GatewayOrderID string    `json:"gateway_order_id" struct:"gateway_order_id"`
	Reference      string    `json:"reference" struct:"reference"`
	Locked         bool      `json:"locked" struct:"locked"`
	Cleared        bool      `json:"cleared" struct:"cleared"`
	Currency       string    `json:"currency" struct:"currency"`
	Channel        string    `json:"channel,omitempty" struct:"channel,omitempty"`
	Amount         int64     `json:"amount" struct:"amount"`
	ForceClear     bool      `json:"force_clear" struct:"force_clear"`
	Status         string    `json:"status" struct:"status"`
	KsherID        string    `json:"ksher_id" struct:"ksher_id"`
	ApiName        string    `json:"api_name" struct:"api_name"`
	Mid            string    `json:"mid" struct:"mid"`
	Signature      string    `json:"signature" struct:"signature"`
	ErrorCode      string    `json:"error_code" struct:"error_code"`
	Acquirer       string    `json:"acquirer" struct:"acquirer"`
	Note           string    `json:"note" struct:"note"`
	Reserved1      *string   `json:"reserved1,omitempty" struct:"reserved1,omitempty"`
	Reserved2      *string   `json:"reserved2,omitempty" struct:"reserved2,omitempty"`
	Reserved3      *string   `json:"reserved3,omitempty" struct:"reserved1,omitempty"`
	Reserved4      *string   `json:"reserved4,omitempty" struct:"reserved1,omitempty"`
	CreatedAt      time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" struct:"updated_at"`
}
