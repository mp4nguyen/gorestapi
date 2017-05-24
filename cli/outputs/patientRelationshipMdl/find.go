package patientRelationshipMdl

import "log"
import "bitbucket.org/restapi/db"

func Find(where string, orderBy string)(patientRelationships PatientRelationships,err error){
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

	response := PatientRelationships{}
	for rows.Next() {
		row := PatientRelationship{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.RelationshipId,&row.PatientId,&row.PersonId,&row.RelationshipType,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.FatherPersonId)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

		response = append(response,&row)
	}

	return response, err
}
