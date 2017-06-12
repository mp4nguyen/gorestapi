package rostersByDateVMdl

import (
	"strconv"

	"bitbucket.org/restapi/models/clinicMdl"
)

func (m *RostersByDateV) FetchClinic() (err error) {
	whereCondition := "clinic_id = " + strconv.Itoa(m.WorkingSiteId)
	tempMapData, err := clinicMdl.MapFind("ClinicId", whereCondition, "clinic_id")
	tempData, ok := tempMapData[strconv.Itoa(m.WorkingSiteId)]
	if ok {
		if len(tempData) > 0 {
			m.Clinic = tempData[0]
		}
	}
	return err
}

func (m *RostersByDateVs) FetchClinic() (err error) {
	foreignKeys := map[string]string{}
	whereCondition := "clinic_id in ("
	for _, row := range *m {
		_, ok := foreignKeys[strconv.Itoa(row.WorkingSiteId)]
		if !ok {
			foreignKeys[strconv.Itoa(row.WorkingSiteId)] = strconv.Itoa(row.WorkingSiteId)
			whereCondition = whereCondition + strconv.Itoa(row.WorkingSiteId) + ","
		}
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := clinicMdl.MapFind("ClinicId", whereCondition, "clinic_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.WorkingSiteId)]
		if ok {
			if len(tempData) > 0 {
				row.Clinic = tempData[0]
			}
		}
	}
	return err
}
