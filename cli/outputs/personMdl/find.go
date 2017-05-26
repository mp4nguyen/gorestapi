package personMdl

import "log"
import "bitbucket.org/restapi/db"

func Find(where string, orderBy string)(persons Persons,err error){
	sqlString := "select person_id,isEnable,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,isPatient,isDoctor,created_by,creation_date,last_updated_by,last_update_date,email,source_id,avatar_id,signature_id,GP_First_name,GP_Last_name,Clinic_Name,GP_Contact,Medicare_No,Medicare_ref,Medicare_Expired from ocs.people"
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
tempMedicareExpired := mysql.NullTime{} 

		rows.Scan(&row.PersonId,&row.IsEnable,&row.Title,&row.FirstName,&row.LastName,&tempDob,&row.Gender,&row.Phone,&row.Mobile,&row.Occupation,&row.Address,&row.SuburbDistrict,&row.Ward,&row.Postcode,&row.StateProvince,&row.Country,&row.IsPatient,&row.IsDoctor,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.Email,&row.SourceId,&row.AvatarId,&row.SignatureId,&row.GPFirstName,&row.GPLastName,&row.ClinicName,&row.GPContact,&row.MedicareNo,&row.MedicareRef,&tempMedicareExpired)
		row.Dob = tempDob.Time 
row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 
row.MedicareExpired = tempMedicareExpired.Time 

		response = append(response,&row)
	}

	return response, err
}
