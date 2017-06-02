package requestMdl

import (
	"database/sql"
	"fmt"
	"time"

	"bitbucket.org/restapi/db"
)

func (inputs Requests) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO requests(request_id,appt_id,patient_id,person_id,type,data,created_by,creation_date,last_updated_by,last_update_date) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
		input.CreationDate = time.Now().UTC()
		input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, input.RequestId, input.ApptId, input.PatientId, input.PersonId, input.Type, input.Data, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"))
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
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
func (input Request) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO requests(request_id,appt_id,patient_id,person_id,type,data,created_by,creation_date,last_updated_by,last_update_date) VALUES "
	vals := []interface{}{}
	input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals, input.RequestId, input.ApptId, input.PatientId, input.PersonId, input.Type, input.Data, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"))
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
