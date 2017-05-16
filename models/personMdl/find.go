package personMdl

import "log"
import "bitbucket.org/restapi/db"

func Find(where string, orderBy string) (persons Persons, err error) {
	sqlString := "select person_id,isEnable,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,isPatient,isDoctor,created_by,creation_date,last_updated_by,last_update_date,email,source_id,avatar_id,signature_id from ocs.people"

	if len(where) > 0 {
		sqlString += (" where " + where)
	}

	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}

	log.Println("personMdl.find.go: will perform sql = ", sqlString)

	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("personMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := Persons{}
	for rows.Next() {
		row := Person{}
		rows.Scan(&row.PersonId, &row.IsEnable, &row.Title, &row.FirstName, &row.LastName, &row.Dob, &row.Gender, &row.Phone, &row.Mobile, &row.Occupation, &row.Address, &row.SuburbDistrict, &row.Ward, &row.Postcode, &row.StateProvince, &row.Country, &row.IsPatient, &row.IsDoctor, &row.CreatedBy, &row.CreationDate, &row.LastUpdatedBy, &row.LastUpdateDate, &row.Email, &row.SourceId, &row.AvatarId, &row.SignatureId)
		response = append(response, row)
	}

	return response, err
}
