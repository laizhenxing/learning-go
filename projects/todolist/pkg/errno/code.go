package errno

var (
	OK = &Errno{Code: 0, Message: "OK"}

	ErrInternalServer = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind           = &Errno{Code: 10002, Message: "Binding request error"}
	ErrParams         = &Errno{Code: 10003, Message: "Params error"}

	ErrDatabase = &Errno{Code: 20001, Message: "Database error"}
	ErrNotFound = &Errno{Code: 20002, Message: "Not found"}
)
