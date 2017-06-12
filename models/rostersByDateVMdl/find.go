package rostersByDateVMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func Find(where string, orderBy string) (rostersByDateVs RostersByDateVs, err error) {
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

	response := RostersByDateVs{}
	for rows.Next() {
		row := RostersByDateV{}
		tempRosterDate := mysql.NullTime{}

		rows.Scan(&row.CompanyId, &row.WorkingSiteId, &row.BookingTypeId, &tempRosterDate)
		row.RosterDate = tempRosterDate.Time

		response = append(response, &row)
	}

	return response, err
}
