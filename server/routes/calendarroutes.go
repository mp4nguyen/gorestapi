package route

import "bitbucket.org/restapi/controllers/calendarCtrl"

var calCtrl = calendarCtrl.CalendarController{}

var calendarroutes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/getCalendar",
		calCtrl.GetCalendar,
	},
}
