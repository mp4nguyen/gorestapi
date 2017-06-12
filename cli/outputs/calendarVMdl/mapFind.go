package calendarVMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *CalendarV, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Kind() == reflect.Int {
		return strconv.Itoa(int(f.Int()))
	} else if f.Kind() == reflect.String {
		return f.String()
	} else {
		return ""
	}
}
func MapFind(groupByField string,where string, orderBy string)(calendarVs map[string]CalendarVs,err error){
	sqlString := "select calendar_id,roster_id,roster_date,clinic_id,booking_type_id,doctor_id,person_id,title,first_name,last_name,gender,from_time,to_time,calendar_date_in_number,calendar_date,calendar_date_in_str,calendar_time_in_str,time_interval,created_by,creation_date,last_updated_by,last_update_date,reserve_id from ocs.calendars_v"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("calendarVMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string]CalendarVs{}
	for rows.Next() {
		row := CalendarV{}
		tempRosterDate := mysql.NullTime{} 
tempFromTime := mysql.NullTime{} 
tempToTime := mysql.NullTime{} 
tempCalendarDate := mysql.NullTime{} 
tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.CalendarId,&row.RosterId,&tempRosterDate,&row.ClinicId,&row.BookingTypeId,&row.DoctorId,&row.PersonId,&row.Title,&row.FirstName,&row.LastName,&row.Gender,&tempFromTime,&tempToTime,&row.CalendarDateInNumber,&tempCalendarDate,&row.CalendarDateInStr,&row.CalendarTimeInStr,&row.TimeInterval,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.ReserveId)
		row.RosterDate = tempRosterDate.Time 
row.FromTime = tempFromTime.Time 
row.ToTime = tempToTime.Time 
row.CalendarDate = tempCalendarDate.Time 
row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, &row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = CalendarVs{&row}
		}
	}

	return response, err
}
