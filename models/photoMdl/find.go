package photoMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func Find(where string, orderBy string) (photos Photos, err error) {
	sqlString := "select photo_id,request_id,appt_id,patient_id,person_id,type,data,uri,created_by,creation_date,last_updated_by,last_update_date from ocs.photos"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("photoMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := Photos{}
	for rows.Next() {
		row := Photo{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}

		rows.Scan(&row.PhotoId, &row.RequestId, &row.ApptId, &row.PatientId, &row.PersonId, &row.Type, &row.Data, &row.Uri, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate)
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time

		response = append(response, &row)
	}

	return response, err
}
