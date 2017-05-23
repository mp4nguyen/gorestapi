package personMdl

import (
	"database/sql"
	"fmt"
	"time"

	"bitbucket.org/restapi/db"
)

func (inputs Persons) Create() (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO people(person_id,isEnable,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,isPatient,isDoctor,created_by,creation_date,last_updated_by,last_update_date,email,source_id,avatar_id,signature_id) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
		input.CreationDate = time.Now()
		input.LastUpdateDate = time.Now()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, input.PersonId, input.IsEnable, input.Title, input.FirstName, input.LastName, input.Dob.Format("2006-01-02 15:04:05"), input.Gender, input.Phone, input.Mobile, input.Occupation, input.Address, input.SuburbDistrict, input.Ward, input.Postcode, input.StateProvince, input.Country, input.IsPatient, input.IsDoctor, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.Email, input.SourceId, input.AvatarId, input.SignatureId)
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
func (input Person) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO people(person_id,isEnable,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,isPatient,isDoctor,created_by,creation_date,last_updated_by,last_update_date,email,source_id,avatar_id,signature_id) VALUES "
	vals := []interface{}{}
	input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals, input.PersonId, input.IsEnable, input.Title, input.FirstName, input.LastName, input.Dob.Format("2006-01-02 15:04:05"), input.Gender, input.Phone, input.Mobile, input.Occupation, input.Address, input.SuburbDistrict, input.Ward, input.Postcode, input.StateProvince, input.Country, input.IsPatient, input.IsDoctor, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.Email, input.SourceId, input.AvatarId, input.SignatureId)

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
