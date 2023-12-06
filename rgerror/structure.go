package rgerror

// PubsubResponseError is throw error from nested function
type PubsubResponseError struct {
	Message       string
	IsAcknowledge bool
}
