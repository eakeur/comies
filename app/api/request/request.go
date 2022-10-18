package request

import (
	"comies/app/core/id"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Request struct {
	*http.Request
}

func (r Request) Param(name string) string {
	return chi.URLParam(r.Request, name)
}

func (r Request) IDParam(name string) (id.ID, error) {
	return id.FromString(r.Param(name))
}

func (r Request) JSONBody(i interface{}) error {
	return json.NewDecoder(r.Body).Decode(i)
}

func (r Request) GetQuery(name string) string {
	return r.URL.Query().Get(name)
}
