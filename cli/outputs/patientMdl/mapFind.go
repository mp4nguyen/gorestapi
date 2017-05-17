package patientMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *Patient, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Kind() == reflect.Int {
		return strconv.Itoa(int(f.Int()))
	} else if f.Kind() == reflect.String {
		return f.String()
	} else {
		return ""
	}
}
func MapFind(groupByField string,where string, orderBy string)(patients map[string][]Patient,err error){
	sqlString := "select patient_id,medical_company_id,user_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,source_id,father_patient_id from ocs.patients"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("patientMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string][]Patient{}
	for rows.Next() {
		row := Patient{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.PatientId,&row.MedicalCompanyId,&row.UserId,&row.PersonId,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.SourceId,&row.FatherPatientId)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = []Patient{row}
		}
	}

	return response, err
}
