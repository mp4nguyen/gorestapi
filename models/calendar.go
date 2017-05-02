package models

import (
	"log"
	"time"

	"bitbucket.org/restapi/db"
)

var layout string = "2006-01-02 15:04:05"

type Calendar struct {
	CalId            int       `json:"calId"`
	RosterId         int       `json:"rosterId"`
	DoctorId         int       `json:"doctorId"`
	DoctorName       string    `json:"doctorName"`
	CalendarFromTime time.Time `json:"calendarFromTime"`
	CalendarToTime   time.Time `json:"calendarFromTime"`
	SiteId           int       `json:"siteId"`
	CalendarDate     string    `json:"calendarDate"`
	CalendarTime     string    `json:"calendarTime"`
}

// CalendarFromTimeInTime time.Time `json:"calendarFromTime"`
// CalendarToTimeInTime   time.Time `json:"calendarFromTime"`

type Calendars struct {
	Calendars []Calendar `json:"calendars"`
}

type CalendarModel struct{}

func (m CalendarModel) GetCalendar(id int, from string, to string) (calendars Calendars, err error) {

	rows, err := db.GetDB().Query("select cal_id,roster_id,doctor_id,doctor_name,calendar_from_time,calendar_to_time,site_id,calendar_date,calendar_time from calendar_tam where enable = 'Y' and site_id = ? and calendar_from_time >= ? and calendar_from_time <= ? order by calendar_date,doctor_id,calendar_from_time", id, from, to)
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

	return Response, err
}
