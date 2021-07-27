package errs

import (
	"errors"
	"net/http"
)

type HttpError interface {
	Error() string
	StatusCode() int
}
type BadRequest struct {
	error
}

func NewBadRequest(err string) *BadRequest {
	return &BadRequest{errors.New(err)}
}
func (e *BadRequest) StatusCode() int {
	return http.StatusBadRequest
}
