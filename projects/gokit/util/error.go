package util

type MyError struct {
	Code    int
	Message string
}

func (m *MyError) Error() string {
	return m.Message
}

func NewError(code int, msg string) error {
	return &MyError{
		Code:    code,
		Message: msg,
	}
}


