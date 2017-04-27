package route

import (
	"bitbucket.org/restapi/controllers"
	"bitbucket.org/restapi/middleware"
	"bitbucket.org/restapi/myjwt"
)

var userCtrl = controllers.UserController{}

var userroutes = Routes{
	Route{
		"AllUsers",
		"GET",
		"/users",
		middleware.AddMiddleware(userCtrl.UsersRetrieve, myjwt.JWTMW),
	},
	Route{
		"OneUser",
		"GET",
		"/users/{id:[0-9]+}",
		middleware.AddMiddleware(userCtrl.GetUser, myjwt.JWTMW),
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		userCtrl.CreateUser,
	},
	Route{
		"Login",
		"POST",
		"/login",
		userCtrl.Login,
	},
	Route{
		"Login",
		"GET",
		"/afterlogin",
		middleware.AddMiddleware(userCtrl.AfterLogin, myjwt.JWTMW),
	},
	Route{
		"Login",
		"GET",
		"/afterlogin2",
		userCtrl.AfterLogin,
	},
}
