/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package patientAppointmentMdl

type PatientAppointment struct{
	Requests []requestMdl.Request `json:"requestss"`
	}

