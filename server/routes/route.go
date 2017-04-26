package route

import (
	"net/http"

	"bitbucket.org/restapi/middleware"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1/").Subrouter()
	for _, route := range userroutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(middleware.Logger(route.HandlerFunc, route.Name))

	}
	return router
}
