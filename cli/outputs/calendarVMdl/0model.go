package calendarVMdl

import "time"

type CalendarV struct{
	CalendarId int `json:"calendarId" mysql:"calendar_id"`
	RosterId int `json:"rosterId" mysql:"roster_id"`
	RosterDate time.Time `json:"rosterDate" mysql:"roster_date"`
	ClinicId int `json:"clinicId" mysql:"clinic_id"`
	BookingTypeId int `json:"bookingTypeId" mysql:"booking_type_id"`
	DoctorId int `json:"doctorId" mysql:"doctor_id"`
	PersonId int `json:"personId" mysql:"person_id"`
	Title string `json:"title" mysql:"title"`
	FirstName string `json:"firstName" mysql:"first_name"`
	LastName string `json:"lastName" mysql:"last_name"`
	Gender string `json:"gender" mysql:"gender"`
	FromTime time.Time `json:"fromTime" mysql:"from_time"`
	ToTime time.Time `json:"toTime" mysql:"to_time"`
	CalendarDateInNumber int `json:"calendarDateInNumber" mysql:"calendar_date_in_number"`
	CalendarDate time.Time `json:"calendarDate" mysql:"calendar_date"`
	CalendarDateInStr char `json:"calendarDateInStr" mysql:"calendar_date_in_str"`
	CalendarTimeInStr char `json:"calendarTimeInStr" mysql:"calendar_time_in_str"`
	TimeInterval int `json:"timeInterval" mysql:"time_interval"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	ReserveId int `json:"reserveId" mysql:"reserve_id"`
	}

type CalendarVs []*CalendarV