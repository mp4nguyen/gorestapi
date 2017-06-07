package patientRelationshipVMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func FindById(id int64) (patientRelationshipVs PatientRelationshipV, err error) {
	rs := db.GetDB().QueryRow("select relationship_id,relationship_type,patient_id,person_id,father_person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,email,avatar_id,avatar_url,signature_id,signature_url,GP_First_name,GP_Last_name,Clinic_Name,GP_Contact,Medicare_No,Medicare_ref,Medicare_Expired from ocs.patient_relationships_v where relationship_id = ?", id)
	if err != nil {
		log.Println("patientRelationshipVMdl.find.go: All() err = ", err)
	}
	row := PatientRelationshipV{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}
	tempDob := mysql.NullTime{}
	tempMedicareExpired := mysql.NullTime{}

	rs.Scan(&row.RelationshipId, &row.RelationshipType, &row.PatientId, &row.PersonId, &row.FatherPersonId, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.Title, &row.FirstName, &row.LastName, &tempDob, &row.Gender, &row.Phone, &row.Mobile, &row.Occupation, &row.Address, &row.SuburbDistrict, &row.Ward, &row.Postcode, &row.StateProvince, &row.Country, &row.Email, &row.AvatarId, &row.AvatarUrl, &row.SignatureId, &row.SignatureUrl, &row.GPFirstName, &row.GPLastName, &row.ClinicName, &row.GPContact, &row.MedicareNo, &row.MedicareRef, &tempMedicareExpired)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time
	row.Dob = tempDob.Time
	row.MedicareExpired = tempMedicareExpired.Time

	return row, err
}
