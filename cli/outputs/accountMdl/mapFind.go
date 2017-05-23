package accountMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *Account, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Kind() == reflect.Int {
		return strconv.Itoa(int(f.Int()))
	} else if f.Kind() == reflect.String {
		return f.String()
	} else {
		return ""
	}
}
func MapFind(groupByField string,where string, orderBy string)(accounts map[string][]Account,err error){
	sqlString := "select password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username from ocs.accounts"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("accountMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string][]Account{}
	for rows.Next() {
		row := Account{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 
tempCreated := mysql.NullTime{} 
tempLastupdated := mysql.NullTime{} 

		rows.Scan(&row.Password,&row.Email,&row.UserType,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.PersonId,&row.DoctorId,&row.PatientId,&row.CompanyId,&row.EmailVerified,&row.Realm,&row.Credentials,&row.Challenges,&row.VerificationToken,&row.Status,&tempCreated,&tempLastupdated,&row.Id,&row.Username)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 
row.Created = tempCreated.Time 
row.Lastupdated = tempLastupdated.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = []Account{row}
		}
	}

	return response, err
}
