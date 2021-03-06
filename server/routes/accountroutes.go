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
	Route{
		"Logout",
		"POST",
		"/logout",
		accountCtrl.Logout,
	},
	Route{
		"CheckAvailableAccount",
		"POST",
		"/checkAvailableAccount",
		accountCtrl.CheckAvailableAccount,
	},
	Route{
		"Signup",
		"POST",
		"/signup",
		accountCtrl.Signup,
	},
	Route{
		"Signup2",
		"POST",
		"/signup2",
		accountCtrl.Signup2,
	},
	Route{
		"NewMember",
		"POST",
		"/newMember",
		accountCtrl.NewMember,
	},
	Route{
		"UpdateMember",
		"POST",
		"/updateMember",
		accountCtrl.UpdateMember,
	},
	Route{
		"ChangePassword",
		"POST",
		"/changePassword",
		accountCtrl.ChangePassword,
	},
}
