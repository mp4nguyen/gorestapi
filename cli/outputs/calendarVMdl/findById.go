package calendarVMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(calendarVs CalendarV,err error){
	rs := db.GetDB().QueryRow("select calendar_id,roster_id,roster_date,clinic_id,booking_type_id,doctor_id,person_id,title,first_name,last_name,gender,from_time,to_time,calendar_date_in_number,calendar_date,calendar_date_in_str,calendar_time_in_str,time_interval,created_by,creation_date,last_updated_by,last_update_date,reserve_id from ocs.calendars_v where  = ?",id)
	if err != nil {
		log.Println("calendarVMdl.find.go: All() err = ", err)
	}
	row := CalendarV{}
		tempRosterDate := mysql.NullTime{} 
tempFromTime := mysql.NullTime{} 
tempToTime := mysql.NullTime{} 
tempCalendarDate := mysql.NullTime{} 
tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

	rs.Scan(&row.CalendarId,&row.RosterId,&tempRosterDate,&row.ClinicId,&row.BookingTypeId,&row.DoctorId,&row.PersonId,&row.Title,&row.FirstName,&row.LastName,&row.Gender,&tempFromTime,&tempToTime,&row.CalendarDateInNumber,&tempCalendarDate,&row.CalendarDateInStr,&row.CalendarTimeInStr,&row.TimeInterval,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.ReserveId)
		row.RosterDate = tempRosterDate.Time 
row.FromTime = tempFromTime.Time 
row.ToTime = tempToTime.Time 
row.CalendarDate = tempCalendarDate.Time 
row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

	return row, err
}
