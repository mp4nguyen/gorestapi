package route

import (
	"bitbucket.org/restapi/controllers"
	"bitbucket.org/restapi/myjwt"
)

var userCtrl = controllers.UserController{}

var userroutes = Routes{
	Route{
		"AllUsers",
		"GET",
		"/users",
		myjwt.JWTMW(userCtrl.UsersRetrieve),
	},
	Route{
		"OneUser",
		"GET",
		"/users/{id:[0-9]+}",
		myjwt.JWTMW(userCtrl.GetUser),
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
		myjwt.JWTMW(userCtrl.AfterLogin),
	},
}
