package cloudsql

import "time"

type RestgoBank struct {
	ID                          string      `json:"id" struct:"id" gorm:"primaryKey"`
	AccountName                 string      `json:"account_name" struct:"account_name" ` //ชื่อ บช
	BankNameTH                  string      `json:"bank_name_th" struct:"bank_name_th"`
	BankNameEN                  string      `json:"bank_name_en" struct:"bank_name_en"`
	BankLogoImage               string      `json:"bank_logo_image" struct:"bank_logo_image"` //รูป logo ธนาคาร
	BankNumber                  string      `json:"bank_number" struct:"bank_number"`
	BankNumberImageUrl          *string     `json:"bank_number_image_url,omitempty" struct:"bank_number_image_url,omitempty"`
	BankNumberImageStorage      *string     `json:"bank_number_image_storage,omitempty" struct:"bank_number_image_storage,omitempty"`
	PromptpayNumber             *string     `json:"prompt_pay_number,omitempty" struct:"prompt_pay_number,omitempty"`
	PromptpayNumberImageUrl     *string     `json:"prompt_pay_number_image_url,omitempty" struct:"prompt_pay_number_image_url,omitempty"`
	PromptpayNumberImageStorage *string     `json:"prompt_pay_number_image_storage,omitempty" struct:"prompt_pay_number_image_storage,omitempty"`
	TypeOfBankID                string      `json:"type_of_bank_id" struct:"type_of_bank_id" gorm:"foreignKey:type_of_bank_id"`
	TypeOfBank                  *TypeOfBank `json:"type_of_bank,omitempty" struct:"type_of_bank,omitempty" gorm:"foreignKey:type_of_bank_id"`
	TypeOfAccount               string      `json:"type_of_account" struct:"type_of_account"` //ประเภท บช (ออมทรัพย์ / ฝากประจำ)
	IsDefault                   bool        `json:"is_default" struct:"is_default"`
	IsSuspended                 bool        `json:"is_suspended" struct:"is_suspended"`
	CreatedAt                   time.Time   `json:"created_at" struct:"created_at"`
	UpdatedAt                   time.Time   `json:"updated_at" struct:"updated_at"`
	DeletedAt                   *time.Time  `json:"deleted_at,omitempty" structs:"deleted_at,omitempty"`
	EditorID                    *string     `json:"editor_id,omitempty" structs:"editor_id,omitempty"`
	EditorType                  *string     `json:"editor_type,omitempty" structs:"editor_type,omitempty"`
}
