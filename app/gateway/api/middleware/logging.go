package middleware

import (
	"context"
	"net/http"
)

type LoggerContextKey struct{}

func (m Middlewares) Logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handler.ServeHTTP(writer, request.WithContext(
			context.WithValue(
				request.Context(), LoggerContextKey{}, m.managers.Logger.With(
					"method", request.Method, "path", request.URL.Path,
				),
			),
		))
	})

}
