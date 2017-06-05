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
		"GetAppointments",
		"POST",
		"/getAppointments",
		skinRequestCtrl.GetAppointments,
	},
	Route{
		"GetAppointment",
		"POST",
		"/getAppointment",
		skinRequestCtrl.GetAppointment,
	},
	Route{
		"GetPhoto",
		"POST",
		"/getPhoto",
		skinRequestCtrl.GetPhoto,
	},
}
