package rostersByDateVMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(rostersByDateVs RostersByDateV,err error){
	rs := db.GetDB().QueryRow("select company_id,working_site_id,booking_type_id,roster_date from ocs.rosters_by_date_v where  = ?",id)
	if err != nil {
		log.Println("rostersByDateVMdl.find.go: All() err = ", err)
	}
	row := RostersByDateV{}
		tempRosterDate := mysql.NullTime{} 

	rs.Scan(&row.CompanyId,&row.WorkingSiteId,&row.BookingTypeId,&tempRosterDate)
		row.RosterDate = tempRosterDate.Time 

	return row, err
}
