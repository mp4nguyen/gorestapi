package calendarMdl

import "time"

var layout string = "2006-01-02 15:04:05"

type Calendar struct {
	CalId            int       `json:"calId" sql:"cal_id"`
	RosterId         int       `json:"rosterId" sql:"roster_id"`
	DoctorId         int       `json:"doctorId" sql:"doctor_id"`
	DoctorName       string    `json:"doctorName" sql:"doctor_name"`
	CalendarFromTime time.Time `json:"calendarFromTime" sql:"calendar_from_time"`
	CalendarToTime   time.Time `json:"calendarFromTime" sql:"calendar_to_time"`
	SiteId           int       `json:"siteId" sql:"site_id"`
	CalendarDate     string    `json:"calendarDate" sql:"calendar_date"`
	CalendarTime     string    `json:"calendarTime" sql:"calendar_time"`
}

type Calendar2 struct {
	CalId            int       `json:"calId" sql:"cal_id"`
	RosterId         int       `json:"rosterId" sql:"roster_id"`
	DoctorId         int       `json:"doctorId" sql:"doctor_id"`
	DoctorName       string    `json:"doctorName" sql:"doctor_name"`
	CalendarFromTime time.Time `json:"calendarFromTime" sql:"calendar_from_time"`
	CalendarToTime   time.Time `json:"calendarFromTime" sql:"calendar_to_time"`
	SiteId           int       `json:"siteId" sql:"site_id"`
	CalendarDate     string    `json:"calendarDate" sql:"calendar_date"`
	CalendarTime     string    `json:"calendarTime" sql:"calendar_time"`
}

type Calendars struct {
	Calendars []Calendar `json:"calendars"`
}
