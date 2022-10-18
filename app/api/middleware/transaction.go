package middleware

import (
	"comies/app/data/conn"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Pool(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r.WithContext(conn.WithContext(r.Context(), pool)))
		})
	}
}
