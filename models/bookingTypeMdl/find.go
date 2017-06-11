package bookingTypeMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func Find(where string, orderBy string) (bookingTypes BookingTypes, err error) {
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

	response := BookingTypes{}
	for rows.Next() {
		row := BookingType{}
		tempCreationDate := mysql.NullTime{}
		tempLastUpdateDate := mysql.NullTime{}

		rows.Scan(&row.BookingTypeId, &row.BookingTypeName, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.Icon)
		row.CreationDate = tempCreationDate.Time
		row.LastUpdateDate = tempLastUpdateDate.Time

		response = append(response, &row)
	}

	return response, err
}
