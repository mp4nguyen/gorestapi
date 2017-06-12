package calendarVMdl

import (
	"database/sql"
	"fmt"
	"time"

	"bitbucket.org/restapi/db"
)

func (inputs CalendarVs) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO calendars_v(calendar_id,roster_id,roster_date,clinic_id,booking_type_id,doctor_id,person_id,title,first_name,last_name,gender,from_time,to_time,calendar_date_in_number,calendar_date,calendar_date_in_str,calendar_time_in_str,time_interval,created_by,creation_date,last_updated_by,last_update_date,reserve_id) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
		input.CreationDate = time.Now().UTC()
		input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, input.CalendarId, input.RosterId, input.RosterDate.Format("2006-01-02 15:04:05"), input.ClinicId, input.BookingTypeId, input.DoctorId, input.PersonId, input.Title, input.FirstName, input.LastName, input.Gender, input.FromTime.Format("2006-01-02 15:04:05"), input.ToTime.Format("2006-01-02 15:04:05"), input.CalendarDateInNumber, input.CalendarDate.Format("2006-01-02 15:04:05"), input.CalendarDateInStr, input.CalendarTimeInStr, input.TimeInterval, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.ReserveId)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, errStmt := db.GetDB().Prepare(sqlStr)
	if tx != nil {
		stmt, errStmt = tx.Prepare(sqlStr)
	}
	defer stmt.Close()
	if errStmt != nil {
		fmt.Println("errStmt = ", errStmt)
		return 0, 0, errStmt
	}

	res, errInsert := stmt.Exec(vals...)
	if errInsert != nil {
		fmt.Println("errInsert = ", errInsert)
		return 0, 0, errInsert
	}

	rnoOfRows, _ := res.RowsAffected()
	rlastId, _ := res.LastInsertId()
	return rnoOfRows, rlastId, err
}
func (input CalendarV) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO calendars_v(calendar_id,roster_id,roster_date,clinic_id,booking_type_id,doctor_id,person_id,title,first_name,last_name,gender,from_time,to_time,calendar_date_in_number,calendar_date,calendar_date_in_str,calendar_time_in_str,time_interval,created_by,creation_date,last_updated_by,last_update_date,reserve_id) VALUES "
	vals := []interface{}{}
	input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals, input.CalendarId, input.RosterId, input.RosterDate.Format("2006-01-02 15:04:05"), input.ClinicId, input.BookingTypeId, input.DoctorId, input.PersonId, input.Title, input.FirstName, input.LastName, input.Gender, input.FromTime.Format("2006-01-02 15:04:05"), input.ToTime.Format("2006-01-02 15:04:05"), input.CalendarDateInNumber, input.CalendarDate.Format("2006-01-02 15:04:05"), input.CalendarDateInStr, input.CalendarTimeInStr, input.TimeInterval, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.ReserveId)
	stmt, errStmt := db.GetDB().Prepare(sqlStr)
	if tx != nil {
		stmt, errStmt = tx.Prepare(sqlStr)
	}
	defer stmt.Close()
	if errStmt != nil {
		fmt.Println("errStmt = ", errStmt)
		return 0, 0, errStmt
	}

	res, errInsert := stmt.Exec(vals...)
	if errInsert != nil {
		fmt.Println("errInsert = ", errInsert)
		return 0, 0, errInsert
	}

	rnoOfRows, _ := res.RowsAffected()
	rlastId, _ := res.LastInsertId()
	return rnoOfRows, rlastId, err
}
