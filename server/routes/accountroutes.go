package route

import "bitbucket.org/restapi/controllers/accountCtrl"

var accountroutes = Routes{
	Route{
		"FindAccounts",
		"GET",
		"/accounts",
		accountCtrl.Find,
	},
	Route{
		"LoginAT",
		"POST",
		"/loginAT",
		accountCtrl.LoginAT,
	},
}
