package cloudsql

import "time"

type PropertyAddress struct {
	PropertyID  string    `json:"property_id" struct:"property_id" gorm:"foreignKey:property_id"`
	Property    *Property `json:"property,omitempty" struct:"property,omitempty" gorm:"foreignKey:property_id"`
	Address     *string   `json:"address,omitempty" struct:"address,omitempty"`
	SubDistrict string    `json:"sub_district" struct:"sub_district"`
	District    string    `json:"district" struct:"district"`
	Province    string    `json:"province" struct:"province"`
	ZipCode     *string   `json:"zip_code,omitempty" struct:"zip_code,omitempty"`
	Country     string    `json:"country" struct:"country" gorm:"default:ประเทศไทย"`
	NoticePoint *string   `json:"notice_point,omitempty" struct:"notice_point,omitempty"`
	Latitude    *string   `json:"latitude,omitempty" struct:"latitude,omitempty"`
	Longtitude  *string   `json:"longtitude,omitempty" struct:"longtitude,omitempty"`
	CreatedAt   time.Time `json:"created_at" struct:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" struct:"updated_at"`
}
