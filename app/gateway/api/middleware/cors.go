package middleware

import "net/http"

func (m Middlewares) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		header := rw.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "*")
		header.Add("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(rw, req)
	})
}
