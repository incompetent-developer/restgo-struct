package access

type AccessDetail struct {
	AccessUuid     string        `json:"access_uuid"`
	AccountID      string        `json:"account_id"`
	Role           string        `json:"role"`
	PermissionName []interface{} `json:"permission_name"`
}

type AccessEmailDetails struct {
	AccessUuid      string `json:"access_uuid" structs:"access_uuid"`
	UserId          string `json:"user_ID" structs:"user_ID"`
	UserEmail       string `json:"user_email" structs:"user_email"`
	UserPhoneNumber string `json:"user_phone_number" structs:"user_phone_number"`
	Role            string `json:"role" structs:"role"`
	Permission      string `json:"permission" structs:"permission"`
}

type AccessUserDetail struct {
	AccessUuid string `json:"access_uuid" struct:"access_uuid"`
	UserId     string `json:"user_ID" struct:"user_ID"`
	Role       string `json:"role" struct:"role"`
	Permission string `json:"permission" struct:"permission"`
}
