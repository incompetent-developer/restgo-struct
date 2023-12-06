package ksher

type Ksher struct {
	Key            string `json:"key"`
	CredentialPath string `json:"credential_path"`
}

type ResponseKsher struct {
	ID              string `json:"id" struct:"id" gorm:"primaryKey"`
	OrderType       string `json:"order_type,omitempty" struct:"order_type,omitempty"`
	ErrorMessage    string `json:"error_message" struct:"error_message" `
	MerchantOrderID string `json:"merchant_order_id" struct:"merchant_order_id"`
	Timestamp       string `json:"timestamp,omitempty" struct:"timestamp,omitempty"`
	GatewayOrderID  string `json:"gateway_order_id,omitempty" struct:"gateway_order_id,omitempty"`
	Reference       string `json:"reference,omitempty" struct:"reference,omitempty"`
	Locked          bool   `json:"locked,omitempty" struct:"locked,omitempty"`
	Cleared         bool   `json:"cleared,omitempty" struct:"cleared,omitempty"`
	Currency        string `json:"currency,omitempty" struct:"currency,omitempty"`
	Amount          int64  `json:"amount,omitempty" struct:"amount,omitempty"`
	ForceClear      bool   `json:"force_clear,omitempty" struct:"force_clear,omitempty"`
	Status          string `json:"status,omitempty" struct:"status,omitempty"`
	KsherID         string `json:"ksher_id,omitempty" struct:"ksher_id,omitempty"`
	ApiName         string `json:"api_name,omitempty" struct:"api_name,omitempty"`
	Mid             string `json:"mid,omitempty" struct:"mid,omitempty"`
	Signature       string `json:"signature" struct:"signature"`
	ErrorCode       string `json:"error_code" struct:"error_code"`
	Acquirer        string `json:"acquirer,omitempty" struct:"acquirer,omitempty"`
	Note            string `json:"note,omitempty" struct:"note,omitempty"`
}
