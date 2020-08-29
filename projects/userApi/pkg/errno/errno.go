package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Errno
	Err error
}

func New(errno Errno, err error) *Err {
	return &Err{
		Errno: errno,
		Err:   err,
	}
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	fmt.Println(err, "Decode error")
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, InternalServerError.Message
}
