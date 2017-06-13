package route

import "bitbucket.org/restapi/controllers/onlineBookingCtrl"

var onlinebookingroutes = Routes{
	Route{
		"GetBookingTypes",
		"GET",
		"/getBookingTypes",
		onlineBookingCtrl.GetBookingTypes,
	},
	Route{
		"SearchClinics",
		"POST",
		"/searchClinics",
		onlineBookingCtrl.SearchClinics,
	},
	Route{
		"GetSlots",
		"POST",
		"/getSlots",
		onlineBookingCtrl.GetSlots,
	},
}
