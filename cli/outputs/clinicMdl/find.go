package clinicMdl

import "log"
import "bitbucket.org/restapi/db"

func Find(where string, orderBy string)(clinics Clinics,err error){
	sqlString := "select clinic_id,clinic_name,isEnable,company_id,isBookable,isTelehealth,isCalendar,description,address,suburb_district,ward,postcode,state_province,country,created_by,creation_date,last_updated_by,last_update_date,latitude,longitude,icon_base64 from ocs.clinics"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("clinicMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := Clinics{}
	for rows.Next() {
		row := Clinic{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

		rows.Scan(&row.ClinicId,&row.ClinicName,&row.IsEnable,&row.CompanyId,&row.IsBookable,&row.IsTelehealth,&row.IsCalendar,&row.Description,&row.Address,&row.SuburbDistrict,&row.Ward,&row.Postcode,&row.StateProvince,&row.Country,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.Latitude,&row.Longitude,&row.IconBase64)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

		response = append(response,&row)
	}

	return response, err
}
