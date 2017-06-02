package patientAppointmentMdl

import (
	"strconv"

	"bitbucket.org/restapi/models/requestMdl"
)

func (m *PatientAppointment) FetchRequest() (err error) {
	whereCondition := "appt_id = " + strconv.Itoa(m.ApptId)
	tempMapData, err := requestMdl.MapFind("ApptId", whereCondition, "appt_id")
	tempData, ok := tempMapData[strconv.Itoa(m.ApptId)]
	if ok {
		m.Requests = tempData
	}
	return err
}
func (m *PatientAppointments) FetchRequest() (err error) {
	whereCondition := "appt_id in ("
	for _, row := range *m {
		whereCondition = whereCondition + strconv.Itoa(row.ApptId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := requestMdl.MapFind("ApptId", whereCondition, "appt_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.ApptId)]
		if ok {
			row.Requests = tempData
		}
	}
	return err
}
