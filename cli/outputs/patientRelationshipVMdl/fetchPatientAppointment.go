package patientRelationshipVMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *PatientRelationshipV)FetchPatientAppointment()(err error){
%!(EXTRA string=PatientRelationshipV)	whereCondition := "patient_id = " + strconv.Itoa(m.)
	tempMapData, err := patientAppointmentMdl.MapFind("PatientId",whereCondition, "patient_id")
		tempData, ok := tempMapData[strconv.Itoa(m.)]
		if ok {
			m.Appointments = tempData
		}
	return err
}
func (m *PatientRelationshipVs)FetchPatientAppointment()(err error){
%!(EXTRA string=PatientRelationshipV)	whereCondition := "patient_id in ("
	for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := patientAppointmentMdl.MapFind("PatientId",whereCondition, "patient_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.)]
		if ok {
			row.Appointments = tempData
		}
	}
	return err
}
