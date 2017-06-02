package patientAppointmentMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(patientAppointments PatientAppointment,err error){
	rs := db.GetDB().QueryRow("select appt_id,patient_id,calendar_id,require_date,description,appt_type,appt_status,created_by,creation_date,last_updated_by,last_update_date,appt_date,booking_type_id,clinic_id,doctor_id,roster_id,person_id,source_id,duration,patient_person_id,from_time,to_time from ocs.patient_appointments where appt_id = ?",id)
	if err != nil {
		log.Println("patientAppointmentMdl.find.go: All() err = ", err)
	}
	row := PatientAppointment{}
		tempRequireDate := mysql.NullTime{} 
tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 
tempApptDate := mysql.NullTime{} 
tempFromTime := mysql.NullTime{} 
tempToTime := mysql.NullTime{} 

	rs.Scan(&row.ApptId,&row.PatientId,&row.CalendarId,&tempRequireDate,&row.Description,&row.ApptType,&row.ApptStatus,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&tempApptDate,&row.BookingTypeId,&row.ClinicId,&row.DoctorId,&row.RosterId,&row.PersonId,&row.SourceId,&row.Duration,&row.PatientPersonId,&tempFromTime,&tempToTime)
		row.RequireDate = tempRequireDate.Time 
row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 
row.ApptDate = tempApptDate.Time 
row.FromTime = tempFromTime.Time 
row.ToTime = tempToTime.Time 

	return row, err
}
