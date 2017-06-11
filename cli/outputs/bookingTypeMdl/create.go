package bookingTypeMdl

import "log"
import "bitbucket.org/restapi/db"

func (inputs BookingTypes)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO booking_types(booking_type_id,booking_type_name,isEnable,created_by,creation_date,last_updated_by,last_update_date,icon) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?),"
		vals = append(vals,  input.BookingTypeId, input.BookingTypeName, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.Icon)
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
func (input BookingType)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO booking_types(booking_type_id,booking_type_name,isEnable,created_by,creation_date,last_updated_by,last_update_date,icon) VALUES "
	vals := []interface{}{}
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?)"
	vals = append(vals,  input.BookingTypeId, input.BookingTypeName, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.Icon)
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
