package requestMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(requests Request,err error){
	rs := db.GetDB().QueryRow("select request_id,appt_id,patient_id,person_id,type,data,created_by,creation_date,last_updated_by,last_update_date from ocs.requests where request_id = ?",id)
	if err != nil {
		log.Println("requestMdl.find.go: All() err = ", err)
	}
	row := Request{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

	rs.Scan(&row.RequestId,&row.ApptId,&row.PatientId,&row.PersonId,&row.Type,&row.Data,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

	return row, err
}
