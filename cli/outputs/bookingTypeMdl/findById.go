package bookingTypeMdl

import "log"
import "bitbucket.org/restapi/db"

func FindById(id int64)(bookingTypes BookingType,err error){
	rs := db.GetDB().QueryRow("select booking_type_id,booking_type_name,isEnable,created_by,creation_date,last_updated_by,last_update_date,icon from ocs.booking_types where booking_type_id = ?",id)
	if err != nil {
		log.Println("bookingTypeMdl.find.go: All() err = ", err)
	}
	row := BookingType{}
		tempCreationDate := mysql.NullTime{} 
tempLastUpdateDate := mysql.NullTime{} 

	rs.Scan(&row.BookingTypeId,&row.BookingTypeName,&row.IsEnable,&row.CreatedBy,&tempCreationDate,&row.LastUpdatedBy,&tempLastUpdateDate,&row.Icon)
		row.CreationDate = tempCreationDate.Time 
row.LastUpdateDate = tempLastUpdateDate.Time 

	return row, err
}
