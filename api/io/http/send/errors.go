package send

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

var once = sync.Once{}

var InternalServerError = ResponseError{
	Code:    "ERR_INTERNAL_SERVER_ERROR",
	Message: "Ops! An unexpected error happened here. Please try again later",
}

func JSONError(err error, options ...Option) Response {
	var jsonErr *json.UnmarshalTypeError
	var target string

	if errors.As(err, &jsonErr) {
		target = fmt.Sprintf("%s.%s: %s", jsonErr.Struct, jsonErr.Field, jsonErr.Value)
	}

	return Data(http.StatusBadRequest, ResponseError{
		Code:    "INVALID_REQUEST_BODY",
		Message: "The request body could not be parsed. Please verify if it is not offending the endpoint contract",
		Target:  target,
	}, append(options, WithError(err))...)
}

func IDError(err error, options ...Option) Response {
	return Data(http.StatusBadRequest, ResponseError{
		Code:    "INVALID_ID",
		Message: "The id provided is invalid",
	}, append(options, WithError(err))...)
}

var failures = map[error]Response{}

func RegisterDomainErrorBindings(r map[error]Response) {
	once.Do(func() {
		failures = r
	})
}
