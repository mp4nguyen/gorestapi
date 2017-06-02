package photoMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *Photo, field string) string {
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
func MapFind(groupByField string,where string, orderBy string)(photos map[string][]Photo,err error){
	sqlString := "select photo_id,request_id,appt_id,patient_id,person_id,type,data,uri,created_by,creation_date,last_updated_by,last_update_date from ocs.photos"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("photoMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string][]Photo{}
	for rows.Next() {
		row := Photo{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.PhotoId,&row.RequestId,&row.ApptId,&row.PatientId,&row.PersonId,&row.Type,&row.Data,&row.Uri,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = []Photo{row}
		}
	}

	return response, err
}
