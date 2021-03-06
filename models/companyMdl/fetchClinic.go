package companyMdl

import (
	"strconv"

	"bitbucket.org/restapi/models/clinicMdl"
)

func (m *Companys) FetchClinic() (err error) {
	whereCondition := "company_id in ("
	for _, row := range *m {
		whereCondition = whereCondition + strconv.Itoa(row.CompanyId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := clinicMdl.MapFind("CompanyId", whereCondition, "company_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.CompanyId)]
		if ok {
			row.Clinics = tempData
		}
	}
	return err
}
