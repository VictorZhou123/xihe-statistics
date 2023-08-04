package repositories

// ErrorConcurrentUpdating
type ErrorConcurrentUpdating struct {
	error
}

func NewErrorConcurrentUpdating(err error) ErrorConcurrentUpdating {
	return ErrorConcurrentUpdating{err}
}

func IsErrorConcurrentUpdating(err error) bool {
	_, ok := err.(ErrorConcurrentUpdating)

	return ok
}

// ErrorNoData
type ErrorNoData struct {
	error
}

func NewErrorNoData(err error) ErrorNoData {
	return ErrorNoData{err}
}

func IsErrorNoData(err error) bool {
	_, ok := err.(ErrorNoData)

	return ok
}
