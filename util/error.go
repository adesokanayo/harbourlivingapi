package util

const (
	ErrPermissionDenied = "You are not allowed to perform this operation"
)

type CustomError struct {
	errMessage string
	code       int
}

func (c CustomError) Error2(err error, code int) CustomError {
	return CustomError{
		code:       code,
		errMessage: err.Error(),
	}
}

func (c CustomError) Error() string {
	return c.errMessage
}
