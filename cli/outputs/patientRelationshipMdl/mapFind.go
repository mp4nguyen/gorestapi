package patientRelationshipMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *PatientRelationship, field string) string {
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
func MapFind(groupByField string,where string, orderBy string)(patientRelationships map[string][]PatientRelationship,err error){
	sqlString := "select relationship_id,patient_id,person_id,relationship_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,father_person_id from ocs.patient_relationships"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("patientRelationshipMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string][]PatientRelationship{}
	for rows.Next() {
		row := PatientRelationship{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.RelationshipId,&row.PatientId,&row.PersonId,&row.RelationshipType,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.FatherPersonId)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = []PatientRelationship{row}
		}
	}

	return response, err
}
