package patientMdl

import (
	"database/sql"
	"fmt"
	"time"

	"bitbucket.org/restapi/db"
)

func (inputs Patients) Create() (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO patients(patient_id,medical_company_id,user_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,source_id,father_patient_id) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
		input.CreationDate = time.Now().UTC()
		input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, input.PatientId, input.MedicalCompanyId, input.UserId, input.PersonId, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.SourceId, input.FatherPatientId)
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
func (input Patient) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO patients(patient_id,medical_company_id,user_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,source_id,father_patient_id) VALUES "
	vals := []interface{}{}
	input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals, input.PatientId, input.MedicalCompanyId, input.UserId, input.PersonId, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.SourceId, input.FatherPatientId)
	stmt, errStmt := db.GetDB().Prepare(sqlStr)
	if tx != nil {
		stmt, errStmt = tx.Prepare(sqlStr)
	}
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
