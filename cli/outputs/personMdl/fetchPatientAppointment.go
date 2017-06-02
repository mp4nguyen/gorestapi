package personMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *Person)FetchPatientAppointment()(err error){
%!(EXTRA string=Person)	whereCondition := "person_id = " + strconv.Itoa(m.PersonId)
	tempMapData, err := patientAppointmentMdl.MapFind("PersonId",whereCondition, "person_id")
		tempData, ok := tempMapData[strconv.Itoa(m.PersonId)]
		if ok {
			m.Appointments = tempData
		}
	return err
}
func (m *Persons)FetchPatientAppointment()(err error){
%!(EXTRA string=Person)	whereCondition := "person_id in ("
	for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := patientAppointmentMdl.MapFind("PersonId",whereCondition, "person_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
		if ok {
			row.Appointments = tempData
		}
	}
	return err
}
