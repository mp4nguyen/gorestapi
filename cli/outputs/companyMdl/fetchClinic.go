package companyMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *Companys)FetchClinic()(err error){
	foreignKeys := map[string]string{}
	whereCondition := "company_id in ("
	for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.CompanyId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := clinicMdl.MapFind(whereCondition, "company_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.CompanyId)]
		if ok {
			row.Clinics = tempData
		}
	}
	return err
}
