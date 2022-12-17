package middleware

import "net/http"

func CORS(origins string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			header := rw.Header()
			header.Add("Access-Control-Allow-Origin", origins)
			header.Add("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
			header.Add("Access-Control-Allow-Header", "Accept, Content-Type, Authorization, Access-Control-Allow-Header")
			if req.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(rw, req)
		})
	}
}
