package route

import "bitbucket.org/restapi/controllers/skinRequestCtrl"

var skinrequestroutes = Routes{
	Route{
		"UploadPhoto",
		"POST",
		"/uploadPhoto",
		skinRequestCtrl.UploadHandler,
	},
	Route{
		"GetAppointment",
		"POST",
		"/getAppointment",
		skinRequestCtrl.GetAppointment,
	},
}
