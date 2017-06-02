package patientVMdl

import "log"
import "bitbucket.org/restapi/db"

func (inputs PatientVs)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO patients_v(patient_id,medical_company_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,email,avatar_id,avatar_url,signature_id,signature_url) VALUES "
	vals := []interface{}{}
	for _, input := range inputs {
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals,  input.PatientId, input.MedicalCompanyId, input.PersonId, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.Title, input.FirstName, input.LastName, input.Dob.Format("2006-01-02 15:04:05"), input.Gender, input.Phone, input.Mobile, input.Occupation, input.Address, input.SuburbDistrict, input.Ward, input.Postcode, input.StateProvince, input.Country, input.Email, input.AvatarId, input.AvatarUrl, input.SignatureId, input.SignatureUrl)
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
func (input PatientV)Create(tx *sql.Tx) (noOfRows int64, lastId int64,err error) {
	sqlStr := "INSERT INTO patients_v(patient_id,medical_company_id,person_id,isEnable,created_by,creation_date,last_updated_by,last_update_date,title,first_name,last_name,dob,gender,phone,mobile,occupation,address,suburb_district,ward,postcode,state_province,country,email,avatar_id,avatar_url,signature_id,signature_url) VALUES "
	vals := []interface{}{}
			input.CreationDate = time.Now().UTC()
	input.LastUpdateDate = time.Now().UTC()

	sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	vals = append(vals,  input.PatientId, input.MedicalCompanyId, input.PersonId, input.IsEnable, input.CreatedBy, input.CreationDate.Format("2006-01-02 15:04:05"), input.LastUpdatedBy, input.LastUpdateDate.Format("2006-01-02 15:04:05"), input.Title, input.FirstName, input.LastName, input.Dob.Format("2006-01-02 15:04:05"), input.Gender, input.Phone, input.Mobile, input.Occupation, input.Address, input.SuburbDistrict, input.Ward, input.Postcode, input.StateProvince, input.Country, input.Email, input.AvatarId, input.AvatarUrl, input.SignatureId, input.SignatureUrl)
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
