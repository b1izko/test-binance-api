package errors

// Error codes
const (
	CodeGeneralError = iota
	CodeInvalidRequest
	CodeParseError
	CodeSendRequestError
)

// Error model
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New error
func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
