package response

import (
	"comies/app/gateway/api/middleware"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
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

func WithData(code int, data interface{}) Response {
	return Response{Data: data, code: code}
}

func WithError(code int, payload Error) Response {
	return Response{Error: payload, code: code}
}

func (r Response) Err(err error) Response {
	r.inner = err
	return r
}

func (r Response) Write(w http.ResponseWriter, req *http.Request) {
	logger := LoggerFromContext(req.Context()).With("code", r.code)

	var innErr throw.DetailedError
	if r.inner != nil && errors.As(r.inner, &innErr) {
		logger.Desugar().With(zap.Any("errors", innErr.Stacked())).Error("failed request")
	} else if r.inner != nil {
		logger.Error("errors", r.inner.Error())
	} else {
		logger.Info("finished request")
	}

	w.WriteHeader(r.code)
	_ = json.NewEncoder(w).Encode(r)
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(middleware.LoggerContextKey{}).(*zap.SugaredLogger)
}
