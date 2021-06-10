package errors

type GenericError struct {
	err     error
	message string
}

func CreateGenericError(err error, message string) *GenericError {
	return &GenericError{
		err:     err,
		message: message,
	}
}

func (e GenericError) Message() string {
	return e.message
}

func (e GenericError) Error() string {
	return e.err.Error()
}
