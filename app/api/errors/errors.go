package errors

import (
	"fmt"
	"net/http"
	"time"
)

var (
	InternalError   = Error{Status: http.StatusInternalServerError, Code: "internal server error"}
	ValidationError = Error{Status: http.StatusBadRequest, Code: "validation error"}
	BadRequest      = Error{Status: http.StatusBadRequest, Code: "bad request"}
)

type Error struct {
	Status    int       `json:"status"`
	Code      string    `json:"code"`
	Message   *string   `json:"message,omitempty"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Code)
}

func (e Error) SetMessage(message string) Error {
	e.Message = &message

	return e
}

func (e Error) Is(err error) bool {
	t, ok := err.(Error)
	if !ok {
		return false
	}

	return t.Status == e.Status && t.Code == e.Code
}
