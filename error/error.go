package error

import (
	"fmt"
	"net/http"
)

const (
	ServerError   = 500
	BusinessError = 5001
	Success       = 200
)

var codeMessageSting = map[int]string{
	ServerError: "Internal Server Error",
}

type Error interface {
	Error() string
	ErrorType() int
	Code() int
	Msg() string
	HttpCode() int
}

type errorImpl struct {
	errorType int
	code      int
	msg       string // User readable information
	error     error  // Developer debugging information
	httpCode  int    // Http status code
}

func (e *errorImpl) Error() string {
	return e.error.Error()
}

func (e *errorImpl) Code() int {
	return e.code
}

func (e *errorImpl) Msg() string {
	return e.msg
}

func (e *errorImpl) ErrorType() int {
	return e.errorType
}

func (e *errorImpl) HttpCode() int {
	return e.httpCode
}

func NewServerError(code int, msg string, err error) Error {

	if code == 0 {
		code = ServerError
	}

	if msg == "" {
		msg = CodeToMessage(code)
	}
	if msg == "" {
		msg = "unknown error"
	}

	if err != nil {
		err = fmt.Errorf(msg)
	}

	return NewError(code, msg, err, ServerError)
}

func NewBusinessError(code int, msg string, err error) Error {

	if code == 0 {
		code = BusinessError
	}

	if msg == "" {
		msg = CodeToMessage(code)
	}

	if err == nil {
		err = fmt.Errorf(msg)
	}

	return NewError(code, msg, err, BusinessError)
}

func NewError(code int, msg string, err error, errType int) Error {
	e := &errorImpl{
		code:      code,
		error:     err,
		msg:       msg,
		errorType: errType,
	}
	if text := http.StatusText(code); text != "" {
		e.httpCode = code
	} else {
		e.httpCode = http.StatusOK
	}

	return e
}

func CodeToMessage(code int) string {
	if res, found := codeMessageSting[code]; found {
		return res
	}

	if text := http.StatusText(code); text != "" {
		return text
	}

	return ""
}

func InjectCodeMessage(m map[int]string) {
	for k, v := range m {
		codeMessageSting[k] = v
	}
}
