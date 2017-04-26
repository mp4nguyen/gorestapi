package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/myjwt"
	route "bitbucket.org/restapi/server/routes"
)

func PrimaryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {

	db.Init()
	myjwt.InitKeys()
	// routes := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()
	// userCtrl := controllers.UserController{}
	// routes.HandleFunc("/users", middleware.Logger(userCtrl.CreateUser, "createUser")).Methods("POST")
	// routes.HandleFunc("/users", middleware.Logger(userCtrl.UsersRetrieve, "allUser")).Methods("GET")
	// routes.HandleFunc("/users/{id:[0-9]+}", middleware.Logger(userCtrl.GetUser, "oneUser")).Methods("GET")
	// routes.HandleFunc("/testmw", middleware.Logger(middleware.MiddlewareHandler2(middleware.MiddlewareHandler1(PrimaryHandler)), "testmw")).Methods("GET")
	// routes.HandleFunc("/testmw2", middleware.Logger(PrimaryHandler, "testmw")).Methods("GET")
	http.Handle("/", route.NewRouter())
	http.ListenAndServe(":8080", nil)
}
