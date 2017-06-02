package requestMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *Request, field string) string {
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
func MapFind(groupByField string,where string, orderBy string)(requests map[string][]Request,err error){
	sqlString := "select request_id,appt_id,patient_id,person_id,type,data,created_by,creation_date,last_updated_by,last_update_date from ocs.requests"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("requestMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string][]Request{}
	for rows.Next() {
		row := Request{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.RequestId,&row.ApptId,&row.PatientId,&row.PersonId,&row.Type,&row.Data,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = []Request{row}
		}
	}

	return response, err
}
