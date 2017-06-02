package photoMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func FindById(id int64) (photos Photo, err error) {
	rs := db.GetDB().QueryRow("select photo_id,request_id,appt_id,patient_id,person_id,type,data,uri,created_by,creation_date,last_updated_by,last_update_date from ocs.photos where photo_id = ?", id)
	if err != nil {
		log.Println("photoMdl.find.go: All() err = ", err)
	}
	row := Photo{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}

	rs.Scan(&row.PhotoId, &row.RequestId, &row.ApptId, &row.PatientId, &row.PersonId, &row.Type, &row.Data, &row.Uri, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time

	return row, err
}
