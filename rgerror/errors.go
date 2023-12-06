package rgerror

func (err *PubsubResponseError) Error() string {
	return err.Message
}
