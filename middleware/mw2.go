package middleware

import (
	"fmt"
	"net/http"
)

func MiddlewareHandler2(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware2 is running ...!")
		h.ServeHTTP(w, r)
	})
}
