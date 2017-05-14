package route

import "bitbucket.org/restapi/cli/accountCtrl"

var accountroutes = Routes{
	Route{
		"FindAccounts",
		"GET",
		"/accounts",
		accountCtrl.Find,
	},
}
