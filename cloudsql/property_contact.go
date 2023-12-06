package cloudsql

import "time"

type PropertyContact struct {
	ID                  string       `json:"id" struct:"id" gorm:"primaryKey"`
	PropertyID          string       `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property            *Property    `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	Name                *string      `json:"name,omitempty" struct:"name,omitempty"`
	Firstname           *string      `json:"firstname,omitempty" struct:"firstname,omitempty"`
	Lastname            *string      `json:"lastname,omitempty" struct:"lastname,omitempty"`
	TypeOfPhoneID       *string      `json:"type_of_phone_id" struct:"type_of_phone_id" gorm:"foreignKey:type_of_phone_id"`
	TypeOfPhone         *TypeOfPhone `json:"type_of_phone,omitempty" struct:"type_of_phone,omitempty" gorm:"foreignKey:type_of_phone_id"`
	Phone               *string      `json:"phone,omitempty" struct:"phone,omitempty"`
	Email               *string      `json:"email,omitempty" struct:"email,omitempty"`
	TypeOfOfficePhoneID *string      `json:"type_of_office_phone_id,omitempty" struct:"type_of_office_phone_id,omitempty" gorm:"foreignKey:type_of_office_phone_id"`
	TypeOfOfficePhone   *TypeOfPhone `json:"type_of_office_phone,omitempty" struct:"type_of_office_phone,omitempty" gorm:"foreignKey:type_of_office_phone_id"`
	OfficePhone         *string      `json:"office_phone,omitempty" struct:"office_phone,omitempty"`
	Position            string       `json:"position" struct:"position"`
	AdminAccountID      *string      `json:"admin_account_id,omitempty" struct:"admin_account_id,omitempty"`
	AdminAccount        *Account     `json:"admin_account,omitempty" struct:"admin_account,omitempty" gorm:"foreignKey:admin_account_id"`
	CreatedAt           time.Time    `json:"created_at" struct:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at" struct:"updated_at"`
}
