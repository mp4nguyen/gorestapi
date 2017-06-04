/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package patientRelationshipVMdl

type PatientRelationshipV struct{
	Appointments []patientAppointmentMdl.PatientAppointment `json:"appointmentss"`
	}

