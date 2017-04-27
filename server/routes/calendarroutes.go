package route

import (
	"bitbucket.org/restapi/controllers"
)

var calendarCtrl = controllers.CalendarController{}

var calendarroutes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/getCalendar",
		calendarCtrl.GetCalendar,
	},

}
