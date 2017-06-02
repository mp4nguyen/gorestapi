package patientAppointmentMdl

import "log"
import "bitbucket.org/restapi/db"

func (inputs PatientAppointments)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO patient_appointments(appt_id,patient_id,calendar_id,require_date,description,appt_type,appt_status,created_by,creation_date,last_updated_by,last_update_date,appt_date,booking_type_id,clinic_id,doctor_id,roster_id,person_id,source_id,duration,patient_person_id,from_time,to_time) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals,  input.ApptId, input.PatientId, input.CalendarId, input.RequireDate.Format("2006-01-02 15:04:05"), input.Description, input.ApptType, input.ApptStatus, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.ApptDate.Format("2006-01-02 15:04:05"), input.BookingTypeId, input.ClinicId, input.DoctorId, input.RosterId, input.PersonId, input.SourceId, input.Duration, input.PatientPersonId, input.FromTime.Format("2006-01-02 15:04:05"), input.ToTime.Format("2006-01-02 15:04:05"))
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
func (input PatientAppointment)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO patient_appointments(appt_id,patient_id,calendar_id,require_date,description,appt_type,appt_status,created_by,creation_date,last_updated_by,last_update_date,appt_date,booking_type_id,clinic_id,doctor_id,roster_id,person_id,source_id,duration,patient_person_id,from_time,to_time) VALUES "
	vals := []interface{}{}
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals,  input.ApptId, input.PatientId, input.CalendarId, input.RequireDate.Format("2006-01-02 15:04:05"), input.Description, input.ApptType, input.ApptStatus, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.ApptDate.Format("2006-01-02 15:04:05"), input.BookingTypeId, input.ClinicId, input.DoctorId, input.RosterId, input.PersonId, input.SourceId, input.Duration, input.PatientPersonId, input.FromTime.Format("2006-01-02 15:04:05"), input.ToTime.Format("2006-01-02 15:04:05"))
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
