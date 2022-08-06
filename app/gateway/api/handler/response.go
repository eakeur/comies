package handler

import (
	"comies/app/core/throw"
	"comies/app/gateway/api/middleware"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Error struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Target  interface{} `json:"target,omitempty"`
}

type Response struct {
	inner error
	code  int
	Error error       `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (r Response) Err(err error) Response {
	r.inner = err
	return r
}

func (r Response) Write(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(r.code)
	_ = json.NewEncoder(w).Encode(r)

	go func() {
		logger := LoggerFromContext(req.Context()).With("code", r.code)

		var innErr throw.DetailedError
		if r.inner != nil && errors.As(r.inner, &innErr) {
			logger.Desugar().With(zap.Any("errors", innErr.Stacked())).Error("failed request")
			return
		}

		if r.inner != nil {
			logger.Error("errors", r.inner.Error())
			return
		}

		logger.Info("finished request")
	}()

}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func ResponseWithData(code int, data interface{}) Response {
	return Response{Data: data, code: code}
}

func ResponseWithError(code int, payload Error) Response {
	return Response{Error: payload, code: code}
}

func JSONParsingErrorResponse(err error) Response {
	var (
		typed  *json.UnmarshalTypeError
		target string
	)
	if errors.As(err, &typed) {
		target = fmt.Sprintf("%s.%s: %s", typed.Struct, typed.Field, typed.Value)
	}

	return ResponseWithError(http.StatusBadRequest, Error{
		Code:    "INVALID_REQUEST_BODY",
		Message: "The request body could not be parsed. Please verify if it is not offending the endpoint contract",
		Target:  target,
	}).Err(err)
}

func IDParsingErrorResponse(err error) Response {
	return ResponseWithError(http.StatusBadRequest, Error{
		Code: "INVALID_ID", Message: "The id provided is invalid",
	}).Err(err)
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(middleware.LoggerContextKey{}).(*zap.SugaredLogger)
}
