package patientAppointmentMdl

import "time"

type PatientAppointment struct{
	ApptId int `json:"apptId" mysql:"appt_id"`
	PatientId int `json:"patientId" mysql:"patient_id"`
	CalendarId int `json:"calendarId" mysql:"calendar_id"`
	RequireDate time.Time `json:"requireDate" mysql:"require_date"`
	Description string `json:"description" mysql:"description"`
	ApptType string `json:"apptType" mysql:"appt_type"`
	ApptStatus string `json:"apptStatus" mysql:"appt_status"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	ApptDate time.Time `json:"apptDate" mysql:"appt_date"`
	BookingTypeId int `json:"bookingTypeId" mysql:"booking_type_id"`
	ClinicId int `json:"clinicId" mysql:"clinic_id"`
	DoctorId int `json:"doctorId" mysql:"doctor_id"`
	RosterId int `json:"rosterId" mysql:"roster_id"`
	PersonId int `json:"personId" mysql:"person_id"`
	SourceId int `json:"sourceId" mysql:"source_id"`
	Duration int `json:"duration" mysql:"duration"`
	PatientPersonId int `json:"patientPersonId" mysql:"patient_person_id"`
	FromTime time.Time `json:"fromTime" mysql:"from_time"`
	ToTime time.Time `json:"toTime" mysql:"to_time"`
	}

type PatientAppointments []*PatientAppointment