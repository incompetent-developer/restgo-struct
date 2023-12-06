package restgostruct

// ResponseError is response error structure for sent back to client.
type ResponseError struct {
	Code            int    `json:"code"`
	Message         string `json:"message"`
	SystemErrorCode string `json:"system_error_code"`
}
