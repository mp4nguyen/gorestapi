package patientRelationshipMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func FindById(id int64) (patientRelationships PatientRelationship, err error) {
	rs := db.GetDB().QueryRow("select relationship_id,patient_id,person_id,relationship_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,father_person_id from ocs.patient_relationships where relationship_id = ?", id)
	if err != nil {
		log.Println("patientRelationshipMdl.find.go: All() err = ", err)
	}
	row := PatientRelationship{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}

	rs.Scan(&row.RelationshipId, &row.PatientId, &row.PersonId, &row.RelationshipType, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.FatherPersonId)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time

	return row, err
}
