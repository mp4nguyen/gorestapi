package middleware

import "net/http"

// AddMiddleware adds middleware to a Handler
func AddMiddleware(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
