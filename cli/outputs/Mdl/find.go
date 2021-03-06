package Mdl

import "log"
import "bitbucket.org/restapi/db"

func Find(where string, orderBy string)(s s,err error){
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

	response := s{}
	for rows.Next() {
		row := {}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.CompanyId,&row.CompanyName,&row.IsEnable,&row.Address,&row.SuburbDistrict,&row.Ward,&row.Postcode,&row.StateProvince,&row.Country,&row.Description,&row.Policy,&row.ConditionToBook,&row.LogoPath,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

		response = append(response,&row)
	}

	return response, err
}
