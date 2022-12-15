package middleware

import "net/http"

func CORS() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			header := rw.Header()
			header.Add("Access-Control-Allow-Origin", "*")
			header.Add("Access-Control-Allow-Methods", "*")
			header.Add("Access-Control-Allow-Header", "*")
			if req.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(rw, req)
		})
	}
}
