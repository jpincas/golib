package web

import (
	"net/http"
)

func AddHeader(k, v string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(k, v)
			next.ServeHTTP(w, r)
		})
	}
}
