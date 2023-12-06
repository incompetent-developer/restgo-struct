package token

type Token struct {
	BackyardAccessKey  string `json:"backyard_access_key"`
	BackyardRefreshKey string `json:"backyard_refresh_key"`
	EnduserAccessKey   string `json:"enduser_access_key"`
	EnduserRefreshKey  string `json:"enduser_refresh_key"`
	CredentialPath     string `json:"credential_path"`
}

type TokenDetail struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"-"`
	RefreshUuid  string `json:"-"`
	AtExpires    int64  `json:"-"`
	RtExpires    int64  `json:"-"`
	AtExpireUnix int64  `json:"at_expire_unix"`
	RtExpireUnix int64  `json:"rt_expire_unix"`
}

type TokenUserDetail struct {
	AccessToken  string `json:"access_token" struct:"access_token"`
	RefreshToken string `json:"refresh_token" struct:"refresh_token"`
	AccessUuid   string `json:"-"`
	RefreshUuid  string `json:"-"`
	AtExpires    int64  `json:"-"`
	RtExpires    int64  `json:"-"`
}

type AccessUserDetail struct {
	AccessUuid string `json:"access_uuid" structs:"access_uuid"`
	UserId     string `json:"user_ID" structs:"user_ID"`
	Role       string `json:"role" structs:"role"`
	Permission string `json:"permission" structs:"permission"`
}

type TokenEmailDetails struct {
	AccessToken string `json:"access_token" struct:"access_token"`
	AccessUuid  string `json:"-"`
	AtExpires   int64  `json:"-"`
}

type TokenOtpDetail struct {
	AccessUuid  string `json:"access_uuid" structs:"access_uuid"`
	PhoneNumber string `json:"phone_number" structs:"phone_number"`
	Permission  string `json:"permission" structs:"permission"`
	OtpToken    string `json:"otp_token" structs:"otp_token"`
	RefNo       string `json:"ref_no" structs:"ref_no"`
	Exp         int64  `json:"-"`
}

type AccessEmailDetail struct {
	AccessUuid      string `json:"access_uuid" structs:"access_uuid"`
	UserId          string `json:"user_ID" structs:"user_ID"`
	UserEmail       string `json:"user_email" structs:"user_email"`
	UserPhoneNumber string `json:"user_phone_number" structs:"user_phone_number"`
	Role            string `json:"role" structs:"role"`
	Permission      string `json:"permission" structs:"permission"`
}

type DeleteUserAccountEmailDetail struct {
	AccessUuid string `json:"access_uuid" structs:"access_uuid"`
	UserID     string `json:"user_id" structs:"user_id"`
	Role       string `json:"role" structs:"role"`
	Permission string `json:"permission" structs:"permission"`
}
