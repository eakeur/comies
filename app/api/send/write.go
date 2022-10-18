package send

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var ws = websocket.Upgrader{}

func Write(wr http.ResponseWriter, r Response) {
	code, header, err := r.code, wr.Header(), error(nil)

	for k, v := range r.header {
		header.Set(k, v)
	}

	if r.data != nil {
		err = encodeJSON(wr, r.data)
		if err != nil {
			code = 500
		}
	}

	wr.WriteHeader(code)
}
