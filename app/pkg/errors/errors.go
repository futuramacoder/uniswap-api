package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorNotFound        = New("not found")
	ErrorInternal        = New("internal error")
	ErrorBadRequest      = New("bad request")
	ErrorTooManyRequests = New("too many requests")
)

func New(code string) Error {
	return Error{
		code: code,
	}
}

// Error is a client error type.
type Error struct {
	code         string
	message      string
	err          error
	requestURL   string
	requestBody  []byte
	responseCode int
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Error() string {
	if len(e.message) != 0 && e.err != nil {
		return fmt.Sprintf("%s: %s", e.message, e.err)
	}

	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.code, e.err)
	}

	if len(e.message) != 0 {
		return fmt.Sprintf("%s: %s", e.code, e.message)
	}

	return e.code
}

func (e Error) Message() string {
	return e.message
}

func (e Error) SetMessage(message string) Error {
	e.message = message

	return e
}

func (e Error) RequestURL() string {
	return e.requestURL
}

func (e Error) SetRequestURL(url string) Error {
	e.requestURL = url

	return e
}

func (e Error) RequestBody() []byte {
	return e.requestBody
}

func (e Error) SetRequestBody(body []byte) Error {
	e.requestBody = body

	return e
}

func (e Error) ResponseCode() int {
	return e.responseCode
}

func (e Error) SetResponseCode(code int) Error {
	e.responseCode = code

	return e
}

func (e Error) Wrap(err error) Error {
	e.err = err

	return e
}

func (e Error) Unwrap() error {
	return e.err
}

func (e Error) Is(err error) bool {
	var aserr Error
	if errors.As(err, &aserr) {
		return aserr.code == e.code
	}

	return false
}
