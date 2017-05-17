package patientMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(patients Patient,err error){
	rs := db.GetDB().QueryRow("select patient_id,medical_company_id,user_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,source_id,father_patient_id from ocs.patients where patient_id = ?",id)
	if err != nil {
		log.Println("patientMdl.find.go: All() err = ", err)
	}
	row := Patient{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

	rs.Scan(&row.PatientId,&row.MedicalCompanyId,&row.UserId,&row.PersonId,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.SourceId,&row.FatherPatientId)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

	return row, err
}
