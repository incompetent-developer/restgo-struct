package line

import "net/http"

// Client is account of line
type Client struct {
	Host            string `json:"host"`
	SecureKey       string `json:"secure_key"`
	CloudSecretName string `json:"cloud_secret_name"`
}

// NotifyRequest structure
type NotifyRequest struct {
	Message              string    `json:"message,omitempty"`              //String
	ImageThumbnail       string    `json:"imageThumbnail,omitempty"`       //HTTP/HTTPS URL
	ImageFullsize        string    `json:"imageFullsize,omitempty"`        //HTTP/HTTPS URL
	ImageFile            http.File `json:"imageFile,omitempty"`            //File
	StickerPackageID     int       `json:"stickerPackageId,omitempty"`     //Number
	StickerID            int       `json:"stickerId,omitempty"`            //Number
	NotificationDisabled bool      `json:"notificationDisabled,omitempty"` //Boolean
}

// LineFollowersRequest structure
type LineFollowersRequest struct {
	Limit  string `json:"limit"`
	Strart string `json:"start"`
}

// LineFollowersResponse structure
type LineFollowersResponse struct {
	UserIDs []string `json:"userIds"`
	Next    *string  `json:"next,omitempty"`
}

type LineUserProfileResponse struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureUrl  string `json:"pictureUrl"`
	Language    string `json:"language"`
}

// LineOA structure
type LineOA struct {
	APIHost            string `json:"api_host"`
	APIVersion         string `json:"api_version"`
	ChannelAccessToken string `json:"channel_access_token"`
	CloudSecretName    string `json:"cloud_secret_name"`
}
