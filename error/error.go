package error

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	ServerError   = 500
	BusinessError = 5001
	Success       = 200
)

var codeMessageSting = map[int]string{
	ServerError:   "Internal Server Error",
	BusinessError: "An unexpected error occurred",
}

type Error interface {
	Error() string
	ErrorType() int
	Code() int
	Msg() string
	HttpCode() int
	String() string
	Encode() error
}

type errorImpl struct {
	ErrType     int    `json:"err_type"`
	ErrCode     int    `json:"err_code"`
	ErrMsg      string `json:"err_msg"`       // User readable information
	ErrDetail   string `json:"err_detail"`    // Developer debugging information
	ErrHttpCode int    `json:"err_http_code"` // Http status code
}

func (e *errorImpl) Error() string {
	return e.ErrDetail
}

func (e *errorImpl) Code() int {
	return e.ErrCode
}

func (e *errorImpl) Msg() string {
	return e.ErrMsg
}

func (e *errorImpl) ErrorType() int {
	return e.ErrType
}

func (e *errorImpl) HttpCode() int {
	return e.ErrHttpCode
}

func (e *errorImpl) String() string {
	b, err := json.Marshal(e)
	if err != nil {
		return e.Error()
	}
	return string(b)
}

func (e *errorImpl) Encode() error {
	return fmt.Errorf(e.String())
}

func DecodeError(err error) Error {
	if err == nil {
		return nil
	}
	herr := &errorImpl{}
	errText := err.Error()
	if strings.Contains(errText, "rpc error") && strings.Contains(errText, "err_msg") && strings.Contains(errText, "err_type") {
		i := strings.Index(errText, "{")
		j := strings.LastIndex(errText, "}")
		if j > i && i > 0 && j+1 == len(errText) {
			errText = errText[i : j+1]
		}
	}

	if marshalErr := json.Unmarshal([]byte(errText), &herr); marshalErr != nil {
		unquoteErrText, _ := strconv.Unquote("\"" + errText + "\"")
		if marshalErr = json.Unmarshal([]byte(unquoteErrText), &herr); marshalErr != nil {
			herr = &errorImpl{ErrCode: ServerError, ErrType: ServerError, ErrHttpCode: http.StatusInternalServerError, ErrMsg: CodeToMessage(ServerError), ErrDetail: errText}
		}
	}
	return herr
}

func NewServerError(code int, msg string, err error) Error {
	if code == 0 {
		code = ServerError
	}

	if msg == "" {
		msg = CodeToMessage(code)
	}
	if msg == "" {
		msg = fmt.Sprintf("unknown error, code %d", code)
	}

	if err == nil {
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
		ErrCode:   code,
		ErrDetail: "",
		ErrMsg:    msg,
		ErrType:   errType,
	}

	if err != nil {
		e.ErrDetail = err.Error()
	}
	if text := http.StatusText(code); text != "" {
		e.ErrHttpCode = code
	} else {
		e.ErrHttpCode = http.StatusOK
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
