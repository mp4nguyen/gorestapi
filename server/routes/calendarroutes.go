package route

import "bitbucket.org/restapi/controllers/calendarCtrl"

var calendarroutes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/getCalendar",
		calendarCtrl.GetCalendar,
	},
}
