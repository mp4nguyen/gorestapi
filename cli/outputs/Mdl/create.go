package Mdl

import "log"
import "bitbucket.org/restapi/db"

func (inputs s)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO companies(company_id,company_name,isEnable,address,suburb_district,ward,postcode,state_province,country,description,policy,condition_to_book,logo_path,created_by,creation_date,last_updated_by,last_update_date) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals,  input.CompanyId, input.CompanyName, input.IsEnable, input.Address, input.SuburbDistrict, input.Ward, input.Postcode, input.StateProvince, input.Country, input.Description, input.Policy, input.ConditionToBook, input.LogoPath, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"))
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
func (input )Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO companies(company_id,company_name,isEnable,address,suburb_district,ward,postcode,state_province,country,description,policy,condition_to_book,logo_path,created_by,creation_date,last_updated_by,last_update_date) VALUES "
	vals := []interface{}{}
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals,  input.CompanyId, input.CompanyName, input.IsEnable, input.Address, input.SuburbDistrict, input.Ward, input.Postcode, input.StateProvince, input.Country, input.Description, input.Policy, input.ConditionToBook, input.LogoPath, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"))
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
