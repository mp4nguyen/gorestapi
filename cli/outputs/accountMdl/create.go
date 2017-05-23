package accountMdl

import (
	"fmt"
	"time"

	"bitbucket.org/restapi/db"
)

func (inputs Accounts) Create() (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO accounts(password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
		input.CreationDate = time.Now().UTC()
		input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, input.Password, input.Email, input.UserType, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.PersonId, input.DoctorId, input.PatientId, input.CompanyId, input.EmailVerified, input.Realm, input.Credentials, input.Challenges, input.VerificationToken, input.Status, input.Created.Format("2006-01-02 15:04:05"), input.Lastupdated.Format("2006-01-02 15:04:05"), input.Id, input.Username)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, errStmt := db.GetDB().Prepare(sqlStr)
	defer stmt.Close()
	if errStmt != nil {
		fmt.Println("errStmt = ", errStmt)
		return 0, 0, errStmt
	}

	res, errInsert := stmt.Exec(vals...)
	if errInsert != nil {
		fmt.Println("errInsert = ", errInsert)
		return 0, 0, errInsert
	}

	rnoOfRows, _ := res.RowsAffected()
	rlastId, _ := res.LastInsertId()
	return rnoOfRows, rlastId, err
}
func (input Account) Create() (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO accounts(password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username) VALUES "
	vals := []interface{}{}
	input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals, input.Password, input.Email, input.UserType, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.PersonId, input.DoctorId, input.PatientId, input.CompanyId, input.EmailVerified, input.Realm, input.Credentials, input.Challenges, input.VerificationToken, input.Status, input.Created.Format("2006-01-02 15:04:05"), input.Lastupdated.Format("2006-01-02 15:04:05"), input.Id, input.Username)
	stmt, errStmt := db.GetDB().Prepare(sqlStr)
	defer stmt.Close()
	if errStmt != nil {
		fmt.Println("errStmt = ", errStmt)
		return 0, 0, errStmt
	}

	res, errInsert := stmt.Exec(vals...)
	if errInsert != nil {
		fmt.Println("errInsert = ", errInsert)
		return 0, 0, errInsert
	}

	rnoOfRows, _ := res.RowsAffected()
	rlastId, _ := res.LastInsertId()
	return rnoOfRows, rlastId, err
}
