package send

import (
	"net/http"
)

func Write(wr http.ResponseWriter, r Response) (err error) {
	code, header := r.Code, wr.Header()
	wr.WriteHeader(code)
	for k, v := range r.Header {
		header.Set(k, v)
	}

	if r.Data != nil {
		err = encodeJSON(wr, r.Data)
	}

	return
}
