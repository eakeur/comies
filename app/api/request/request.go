package request

import (
	"comies/app/core/types"
	"comies/app/gateway/data/postgres/conn"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Request struct {
	*http.Request
}

func (r Request) Param(name string) string {
	return chi.URLParam(r.Request, name)
}

func (r Request) IDParam(name string) (types.ID, error) {
	return types.FromString(r.Param(name))
}

func (r Request) JSONBody(i interface{}) error {
	return json.NewDecoder(r.Body).Decode(i)
}

func (r Request) GetQuery(name string) string {
	return r.URL.Query().Get(name)
}

func (r Request) Commit(ctx context.Context) {
	tx, err := conn.TXFromContext(ctx)
	if err == nil {
		tx.Commit(ctx)
	}
}
