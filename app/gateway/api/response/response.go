package response

import (
	"encoding/json"
	"fmt"
	"log"
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

func (r Response) Write(w http.ResponseWriter) {
	if r.inner != nil {
		log.Println(r.inner)
	}

	w.WriteHeader(r.code)
	_ = json.NewEncoder(w).Encode(r)
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
