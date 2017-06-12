package clinicMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *Clinic)FetchCalendarV()(err error){
%!(EXTRA string=Clinic)	whereCondition := "clinic_id = " + strconv.Itoa(m.ClinicId)
	tempMapData, err := calendarVMdl.MapFind("ClinicId",whereCondition, "clinic_id")
		tempData, ok := tempMapData[strconv.Itoa(m.ClinicId)]
		if ok {
			m.Slots = tempData
		}
	return err
}
func (m *Clinics)FetchCalendarV()(err error){
%!(EXTRA string=Clinic)	whereCondition := "clinic_id in ("
	for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.ClinicId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := calendarVMdl.MapFind("ClinicId",whereCondition, "clinic_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.ClinicId)]
		if ok {
			row.Slots = tempData
		}
	}
	return err
}
