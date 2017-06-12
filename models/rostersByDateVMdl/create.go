package rostersByDateVMdl

import (
	"database/sql"
	"fmt"

	"bitbucket.org/restapi/db"
)

func (inputs RostersByDateVs) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO rosters_by_date_v(company_id,working_site_id,booking_type_id,roster_date) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {

		sqlStr += "(?,?,?,?),"
		vals = append(vals, input.CompanyId, input.WorkingSiteId, input.BookingTypeId, input.RosterDate.Format("2006-01-02 15:04:05"))
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
func (input RostersByDateV) Create(tx *sql.Tx) (noOfRows int64, lastId int64, err error) {
	sqlStr := "INSERT INTO rosters_by_date_v(company_id,working_site_id,booking_type_id,roster_date) VALUES "
	vals := []interface{}{}

	sqlStr += "(?,?,?,?)"
	vals = append(vals, input.CompanyId, input.WorkingSiteId, input.BookingTypeId, input.RosterDate.Format("2006-01-02 15:04:05"))
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
