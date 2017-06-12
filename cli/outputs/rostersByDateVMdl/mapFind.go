package rostersByDateVMdl

import "log"
import "bitbucket.org/restapi/db"

func getField(v *RostersByDateV, field string) string {
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
func MapFind(groupByField string,where string, orderBy string)(rostersByDateVs map[string]RostersByDateVs,err error){
	sqlString := "select company_id,working_site_id,booking_type_id,roster_date from ocs.rosters_by_date_v"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("rostersByDateVMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string]RostersByDateVs{}
	for rows.Next() {
		row := RostersByDateV{}
		tempRosterDate := mysql.NullTime{} 

		rows.Scan(&row.CompanyId,&row.WorkingSiteId,&row.BookingTypeId,&tempRosterDate)
		row.RosterDate = tempRosterDate.Time 


		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, &row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = RostersByDateVs{&row}
		}
	}

	return response, err
}
