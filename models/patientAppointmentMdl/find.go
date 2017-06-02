package patientAppointmentMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func Find(where string, orderBy string) (patientAppointments PatientAppointments, err error) {
	sqlString := "select appt_id,patient_id,calendar_id,require_date,description,appt_type,appt_status,created_by,creation_date,last_updated_by,last_update_date,appt_date,booking_type_id,clinic_id,doctor_id,roster_id,person_id,source_id,duration,patient_person_id,from_time,to_time from ocs.patient_appointments"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("patientAppointmentMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := PatientAppointments{}
	for rows.Next() {
		row := PatientAppointment{}
		tempRequireDate := mysql.NullTime{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}
		tempApptDate := mysql.NullTime{}
		tempFromTime := mysql.NullTime{}
		tempToTime := mysql.NullTime{}

		rows.Scan(&row.ApptId, &row.PatientId, &row.CalendarId, &tempRequireDate, &row.Description, &row.ApptType, &row.ApptStatus, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &tempApptDate, &row.BookingTypeId, &row.ClinicId, &row.DoctorId, &row.RosterId, &row.PersonId, &row.SourceId, &row.Duration, &row.PatientPersonId, &tempFromTime, &tempToTime)
		row.RequireDate = tempRequireDate.Time
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time
		row.ApptDate = tempApptDate.Time
		row.FromTime = tempFromTime.Time
		row.ToTime = tempToTime.Time

		response = append(response, &row)
	}

	return response, err
}
