package clinicMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func FindById(id int64) (clinics Clinic, err error) {
	rs := db.GetDB().QueryRow("select clinic_id,clinic_name,isEnable,company_id,isBookable,isTelehealth,isCalendar,description,address,suburb_district,ward,postcode,state_province,country,created_by,creation_date,last_updated_by,last_update_date,latitude,longitude,icon_base64 from ocs.clinics where clinic_id = ?", id)
	if err != nil {
		log.Println("clinicMdl.find.go: All() err = ", err)
	}
	row := Clinic{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}

	rs.Scan(&row.ClinicId, &row.ClinicName, &row.IsEnable, &row.CompanyId, &row.IsBookable, &row.IsTelehealth, &row.IsCalendar, &row.Description, &row.Address, &row.SuburbDistrict, &row.Ward, &row.Postcode, &row.StateProvince, &row.Country, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.Latitude, &row.Longitude, &row.IconBase64)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time

	return row, err
}
