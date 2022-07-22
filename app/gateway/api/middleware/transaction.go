package middleware

import (
	"net/http"
)

func (m Middlewares) Transaction(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := m.managers.Transactions.Begin(request.Context())
		defer m.managers.Transactions.End(ctx)

		handler.ServeHTTP(writer, request.WithContext(ctx))
	})

}
