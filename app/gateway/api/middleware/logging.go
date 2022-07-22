package middleware

import (
	"log"
	"net/http"
	"time"
)

func (m Middlewares) Logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		at := time.Now()

		handler.ServeHTTP(writer, request)

		log.Printf("[%s] Route:%s\tDuration:%s\n",
			request.Method,
			request.URL.Path,
			time.Since(at))
	})

}
