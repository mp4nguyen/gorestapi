package accountMdl

import (
	"log"
	"strconv"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func MapFind(where string, orderBy string) (accounts map[string]Account, err error) {
	sqlString := "select password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username from ocs.accounts"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("accountMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string]Account{}
	for rows.Next() {
		row := Account{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}
		tempCreated := mysql.NullTime{}
		tempLastupdated := mysql.NullTime{}

		rows.Scan(&row.Password, &row.Email, &row.UserType, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.PersonId, &row.DoctorId, &row.PatientId, &row.CompanyId, &row.EmailVerified, &row.Realm, &row.Credentials, &row.Challenges, &row.VerificationToken, &row.Status, &tempCreated, &tempLastupdated, &row.Id, &row.Username)
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time
		row.Created = tempCreated.Time
		row.Lastupdated = tempLastupdated.Time

		response[strconv.Itoa(row.Id)] = row
	}

	return response, err
}
