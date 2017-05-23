package personMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(persons Person,err error){
	rs := db.GetDB().QueryRow("select person_id,isEnable,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,isPatient,isDoctor,created_by,creation_date,last_updated_by,last_update_date,email,source_id,avatar_id,signature_id from ocs.people where person_id = ?",id)
	if err != nil {
		log.Println("personMdl.find.go: All() err = ", err)
	}
	row := Person{}
		tempDob := mysql.NullTime{} 
tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

	rs.Scan(&row.PersonId,&row.IsEnable,&row.Title,&row.FirstName,&row.LastName,&tempDob,&row.Gender,&row.Phone,&row.Mobile,&row.Occupation,&row.Address,&row.SuburbDistrict,&row.Ward,&row.Postcode,&row.StateProvince,&row.Country,&row.IsPatient,&row.IsDoctor,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.Email,&row.SourceId,&row.AvatarId,&row.SignatureId)
		row.Dob = tempDob.Time 
row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

	return row, err
}
