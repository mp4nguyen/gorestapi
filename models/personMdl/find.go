package personMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func Find(where string, orderBy string) (persons Persons, err error) {
	sqlString := "select person_id,isEnable,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,isPatient,isDoctor,created_by,creation_date,last_updated_by,last_update_date,email,source_id,avatar_id,signature_id from ocs.people"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("personMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := Persons{}
	for rows.Next() {
		row := Person{}
		tempDob := mysql.NullTime{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}

		rows.Scan(&row.PersonId, &row.IsEnable, &row.Title, &row.FirstName, &row.LastName, &tempDob, &row.Gender, &row.Phone, &row.Mobile, &row.Occupation, &row.Address, &row.SuburbDistrict, &row.Ward, &row.Postcode, &row.StateProvince, &row.Country, &row.IsPatient, &row.IsDoctor, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.Email, &row.SourceId, &row.AvatarId, &row.SignatureId)
		row.Dob = tempDob.Time
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time

		response = append(response, &row)
	}

	return response, err
}
