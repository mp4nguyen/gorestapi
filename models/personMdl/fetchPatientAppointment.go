package personMdl

import (
	"fmt"
	"strconv"

	"bitbucket.org/restapi/models/patientAppointmentMdl"
)

func (m *Person) FetchPatientAppointment() (err error) {
	whereCondition := "patient_person_id = " + strconv.Itoa(m.PersonId)
	tempMapData, err := patientAppointmentMdl.MapFind("PatientPersonId", whereCondition, "patient_person_id")
	tempData, ok := tempMapData[strconv.Itoa(m.PersonId)]

	if ok {
		m.Appointments = tempData
	}
	fmt.Println("=========> person = ", m)
	return err
}
func (m *Persons) FetchPatientAppointment() (err error) {
	whereCondition := "person_id in ("
	for _, row := range *m {
		whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := patientAppointmentMdl.MapFind("PersonId", whereCondition, "person_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
		if ok {
			row.Appointments = tempData
		}
	}
	return err
}
