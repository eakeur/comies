package middleware

import (
	"comies/app/io/data/postgres/conn"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TX(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			tx, err := pool.Begin(ctx)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer tx.Rollback(ctx)

			handler.ServeHTTP(w, r.WithContext(conn.WithContext(ctx, tx)))
		})
	}
}

func Pool(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r.WithContext(conn.WithContext(r.Context(), pool)))
		})
	}
}
