package calendarMdl

import (
	"log"
	"time"

	"bitbucket.org/restapi/db"
)

func GetCalendar(id int, from string, to string) (calendars Calendars, err error) {

	start := time.Now()

	rows, err := db.GetDB().Query("select cal_id,roster_id,doctor_id,doctor_name,calendar_from_time,calendar_to_time,site_id,calendar_date,calendar_time from calendar2_v where enable = 'Y' and site_id = ? and calendar_from_time >= ? and calendar_from_time <= ? order by calendar_date,doctor_id,calendar_from_time", id, from, to)
	//rows, err := db.GetDB().Query("select cal_id,roster_id,doctor_id,doctor_name,calendar_from_time,calendar_to_time,site_id,calendar_date,calendar_time from calendar2 ")
	if err != nil {
		log.Println("users.go: All() err = ", err)
	}

	Response := Calendars{}

	for rows.Next() {

		calendar := Calendar{}
		rows.Scan(&calendar.CalId, &calendar.RosterId, &calendar.DoctorId, &calendar.DoctorName, &calendar.CalendarFromTime, &calendar.CalendarToTime, &calendar.SiteId, &calendar.CalendarDate, &calendar.CalendarTime)
		//calendar.CalendarFromTimeInTime,err := time.Parse(layout, calendar.CalendarFromTime)
		Response.Calendars = append(Response.Calendars, calendar)
	}

	log.Printf("sql with normal way duration = %s", time.Since(start))

	return Response, err
}
