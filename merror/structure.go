package merror

// HTTPResponseError is throw error from nested function
type HTTPResponseError struct {
	Message         string
	StatusCode      int
	SystemErrorCode string
}

// PubsubResponseError is throw error from nested function
type PubsubResponseError struct {
	Message       string
	IsAcknowledge bool
}
