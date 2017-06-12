package rostersByDateVMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *RostersByDateV)FetchClinic()(err error){
%!(EXTRA string=RostersByDateV)	whereCondition := "clinic_id = " + strconv.Itoa(m.ClinicId) 
	tempMapData, err := clinicMdl.MapFind("ClinicId",whereCondition, "clinic_id")
		tempData, ok := tempMapData[strconv.Itoa(m.ClinicId)]
		if ok {
			if len(tempData) > 0 {
			m.Clinic = tempData[0]
			}
		}
	return err
}
func (m *RostersByDateVs)FetchClinic()(err error){
%!(EXTRA string=RostersByDateV)	foreignKeys := map[string]string{}
	whereCondition := "clinic_id in ("
	for _, row := range *m {
		_, ok := foreignKeys[strconv.Itoa(row.WorkingSiteId)]
		if !ok {
			foreignKeys[strconv.Itoa(row.WorkingSiteId)] = strconv.Itoa(row.WorkingSiteId)
			whereCondition = whereCondition + strconv.Itoa(row.WorkingSiteId) + ","
		}
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := clinicMdl.MapFind("ClinicId",whereCondition, "clinic_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.ClinicId)]
		if ok {
			if len(tempData) > 0 {
			row.Clinic = tempData[0]
			}
		}
	}
	return err
}
