package accountMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func Find() (accounts Accounts, err error) {
	rows, err := db.GetDB().Query("select password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username from ocs.accounts")
	if err != nil {
		log.Println("accountMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := Accounts{}
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

		response = append(response, &row)
	}

	return response, err
}
