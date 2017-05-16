package accountMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64) (accounts Account, err error) {
	rs := db.GetDB().QueryRow("select password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username from ocs.accounts where id = ?", id)
	if err != nil {
		log.Println("accountMdl.find.go: All() err = ", err)
	}
	row := Account{}
	rs.Scan(&row.Password, &row.Email, &row.UserType, &row.IsEnable, &row.CreatedBy, &row.CreationDate, &row.LastUpdatedBy, &row.LastUpdateDate, &row.PersonId, &row.DoctorId, &row.PatientId, &row.CompanyId, &row.EmailVerified, &row.Realm, &row.Credentials, &row.Challenges, &row.VerificationToken, &row.Status, &row.Created, &row.Lastupdated, &row.Id, &row.Username)
	return row, err
}
