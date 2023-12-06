package thaibulk

type Thaibulk struct {
	UrlSendOtp     string `json:"url_send_otp"`
	UrlVerify      string `json:"url_verify"`
	AppKey         string `json:"app_key"`
	AppSecret      string `json:"app_secret"`
	CredentialPath string `json:"credential_path"`
}
