package merror

func (err *HTTPResponseError) Error() string {
	return err.Message
}

func (err *PubsubResponseError) Error() string {
	return err.Message
}
