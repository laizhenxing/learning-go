package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (e *Errno) Error() string {
	return e.Message
}

type Err struct {
	Errno
	Err error
}

func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", e.Code, e.Message, e.Err)
}

func NewErr(errno Errno, err error) *Err {
	return &Err{
		errno,
		err,
	}
}

func DecodeErr(err error) (int, string)  {
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

	return ErrInternalServer.Code, ErrInternalServer.Message
}