package http

import (
	"fmt"
	"net/http"
	"time"
)

func Handle(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	headers := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			Headers(w, r)
			h.ServeHTTP(w, r)
			fmt.Println(fmt.Sprintf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start)))
		}
	}

	h = headers(h)
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func Listen(host string, port int, router http.Handler) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
}
