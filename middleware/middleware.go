package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(ph http.HandlerFunc, middleHandlers ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	fmt.Println("hello?")
	var next http.HandlerFunc = ph
	for _, mw := range middleHandlers {
		fmt.Println("Um?")
		next = mw(ph)
	}
	return next
}
