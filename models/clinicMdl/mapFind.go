package clinicMdl

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func getField(v *Clinic, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	//fmt.Println(" reflect.TypeOf(f).String() = ", f.Kind())
	if f.Kind() == reflect.Int {
		return strconv.Itoa(int(f.Int()))
	} else if f.Kind() == reflect.String {
		return f.String()
	} else {
		return ""
	}
}

func MapFind(groupByField string, where string, orderBy string) (clinics map[string][]Clinic, err error) {
	sqlString := "select clinic_id,clinic_name,isEnable,company_id,isBookable,isTelehealth,isCalendar,description,address,suburb_district,ward,postcode,state_province,country,created_by,creation_date,last_updated_by,last_update_date,latitude,longitude from ocs.clinics"
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

	response := map[string][]Clinic{}
	for rows.Next() {
		row := Clinic{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}

		rows.Scan(&row.ClinicId, &row.ClinicName, &row.IsEnable, &row.CompanyId, &row.IsBookable, &row.IsTelehealth, &row.IsCalendar, &row.Description, &row.Address, &row.SuburbDistrict, &row.Ward, &row.Postcode, &row.StateProvince, &row.Country, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.Latitude, &row.Longitude)
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time
		groupByFieldValue := getField(&row, groupByField) //strconv.Itoa(row.CompanyId)

		//output, _ := json.Marshal(row)
		//fmt.Println("groupByFieldValue = ", groupByFieldValue)
		// fmt.Println(" \nclinics = ", string(output))

		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = []Clinic{row}
		}

	}

	output, _ := json.Marshal(response)
	fmt.Println(" \n\n\nclinics = ", string(output))

	return response, err
}
