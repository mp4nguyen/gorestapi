package accountMdl

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func (m Login) CheckAccount() (isMatch bool, account Account, err error) {

	sqlString := "select password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username from ocs.accounts where username=?"

	fmt.Println(" m.Username = ", m.Username)
	rs := db.GetDB().QueryRow(sqlString, m.Username)

	row := Account{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}
	tempCreated := mysql.NullTime{}
	tempLastupdated := mysql.NullTime{}
	rs.Scan(&row.Password, &row.Email, &row.UserType, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.PersonId, &row.DoctorId, &row.PatientId, &row.CompanyId, &row.EmailVerified, &row.Realm, &row.Credentials, &row.Challenges, &row.VerificationToken, &row.Status, &tempCreated, &tempLastupdated, &row.Id, &row.Username)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time
	row.Created = tempCreated.Time
	row.Lastupdated = tempLastupdated.Time

	if row.Password == "" {
		return false, Account{}, errors.New("Username does not exist")
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(m.Password))

	if errBcrypt != nil {
		return false, Account{}, errors.New("Wrong password")
	} else {
		return true, row, err
	}

}
