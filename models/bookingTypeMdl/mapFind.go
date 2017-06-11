package bookingTypeMdl

import (
	"log"
	"reflect"
	"strconv"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func getField(v *BookingType, field string) string {
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
func MapFind(groupByField string, where string, orderBy string) (bookingTypes map[string]BookingTypes, err error) {
	sqlString := "select booking_type_id,booking_type_name,isEnable,created_by,creation_date,last_updated_by,last_update_date,icon from ocs.booking_types"
	if len(where) > 0 {
		sqlString += (" where " + where)
	}
	if len(orderBy) > 0 {
		sqlString += (" order by " + orderBy)
	}
	rows, err := db.GetDB().Query(sqlString)
	if err != nil {
		log.Println("bookingTypeMdl.find.go: All() err = ", err)
	}
	defer rows.Close()

	response := map[string]BookingTypes{}
	for rows.Next() {
		row := BookingType{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}

		rows.Scan(&row.BookingTypeId, &row.BookingTypeName, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.Icon)
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time

		groupByFieldValue := getField(&row, groupByField)
		group, ok := response[groupByFieldValue]
		if ok {
			group = append(group, &row)
			response[groupByFieldValue] = group
		} else {
			response[groupByFieldValue] = BookingTypes{&row}
		}
	}

	return response, err
}
