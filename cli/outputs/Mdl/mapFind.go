package Mdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *, field string) string {
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
func MapFind(groupByField string,where string, orderBy string)(s map[string]s,err error){
	sqlString := "select company_id,company_name,isEnable,address,suburb_district,ward,postcode,state_province,country,description,policy,condition_to_book,logo_path,created_by,creation_date,last_updated_by,last_update_date from ocs.companies"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("Mdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string]s{}
	for rows.Next() {
		row := {}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.CompanyId,&row.CompanyName,&row.IsEnable,&row.Address,&row.SuburbDistrict,&row.Ward,&row.Postcode,&row.StateProvince,&row.Country,&row.Description,&row.Policy,&row.ConditionToBook,&row.LogoPath,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, &row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = s{&row}
		}
	}

	return response, err
}
