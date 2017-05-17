package route

import (
	"bitbucket.org/restapi/controllers/userCtrl"
	"bitbucket.org/restapi/middleware"
	"bitbucket.org/restapi/myjwt"
)

var userroutes = Routes{
	Route{
		"AllUsers",
		"GET",
		"/users",
		middleware.AddMiddleware(userCtrl.UsersRetrieve, middleware.IsAuthenticatedATMW),
	},
	Route{
		"OneUser",
		"GET",
		"/users/{id:[0-9]+}",
		middleware.AddMiddleware(userCtrl.GetUser, middleware.IsAuthenticatedATMW),
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		userCtrl.CreateUser,
	},
	Route{
		"LoginJWT",
		"POST",
		"/loginJWT",
		userCtrl.LoginJWT,
	},

	Route{
		"Logout",
		"POST",
		"/logout",
		middleware.AddMiddleware(userCtrl.Logout, middleware.IsAuthenticatedATMW),
	},
	Route{
		"afterloginjwt",
		"GET",
		"/afterloginjwt",
		middleware.AddMiddleware(userCtrl.AfterLogin, myjwt.JWTMW),
	},
	Route{
		"afterloginat",
		"GET",
		"/afterloginat",
		middleware.AddMiddleware(userCtrl.AfterLogin, middleware.IsAuthenticatedATMW),
	},
	Route{
		"Login",
		"GET",
		"/afterlogin2",
		userCtrl.AfterLogin,
	},
}
