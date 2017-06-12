package rostersByDateVMdl

import (
	"log"
	"time"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func FindByBookingTypeAndDate(bookingTypeId int, rosterDate time.Time) (rostersByDateVs RostersByDateVs, err error) {
	rows, err := db.GetDB().Query("select company_id,working_site_id,booking_type_id,roster_date from ocs.rosters_by_date_v where  booking_type_id = ? and roster_date = ?", bookingTypeId, rosterDate.Format("2006-01-02"))
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
