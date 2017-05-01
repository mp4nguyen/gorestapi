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

func createRouter(router *mux.Router, routes Routes) {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(middleware.Logger(middleware.CORS(route.HandlerFunc), route.Name))
	}
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1/").Subrouter()
	createRouter(router, userroutes)
	createRouter(router, calendarroutes)
	return router
}
