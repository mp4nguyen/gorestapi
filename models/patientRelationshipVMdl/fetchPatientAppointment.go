package patientRelationshipVMdl

import (
	"strconv"

	"bitbucket.org/restapi/models/patientAppointmentMdl"
)

func (m *PatientRelationshipV) FetchPatientAppointment() (err error) {
	whereCondition := "patient_id = " + strconv.Itoa(m.PatientId)
	tempMapData, err := patientAppointmentMdl.MapFind("PatientId", whereCondition, "patient_id")
	tempData, ok := tempMapData[strconv.Itoa(m.PatientId)]
	if ok {
		m.Appointments = tempData
	}
	return err
}

func (m *PatientRelationshipVs) FetchPatientAppointment() (err error) {
	if len(*m) > 0 {
		whereCondition := "patient_id in ("
		for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.PatientId) + ","
		}
		whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
		tempMapData, err := patientAppointmentMdl.MapFind("PatientId", whereCondition, "patient_id")
		for _, row := range *m {
			tempData, ok := tempMapData[strconv.Itoa(row.PatientId)]
			if ok {
				row.Appointments = tempData
			}
		}
		return err
	} else {
		return nil
	}

}
