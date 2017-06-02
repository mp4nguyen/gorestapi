package patientVMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(patientVs PatientV,err error){
	rs := db.GetDB().QueryRow("select patient_id,medical_company_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,email,avatar_id,avatar_url,signature_id,signature_url from ocs.patients_v where  = ?",id)
	if err != nil {
		log.Println("patientVMdl.find.go: All() err = ", err)
	}
	row := PatientV{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 
tempDob := mysql.NullTime{} 

	rs.Scan(&row.PatientId,&row.MedicalCompanyId,&row.PersonId,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.Title,&row.FirstName,&row.LastName,&tempDob,&row.Gender,&row.Phone,&row.Mobile,&row.Occupation,&row.Address,&row.SuburbDistrict,&row.Ward,&row.Postcode,&row.StateProvince,&row.Country,&row.Email,&row.AvatarId,&row.AvatarUrl,&row.SignatureId,&row.SignatureUrl)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 
row.Dob = tempDob.Time 

	return row, err
}
