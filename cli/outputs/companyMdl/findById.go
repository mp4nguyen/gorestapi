package companyMdl

import (
	"log"

	"bitbucket.org/restapi/db"
	"github.com/go-sql-driver/mysql"
)

func FindById(id int64) (companys Company, err error) {
	rs := db.GetDB().QueryRow("select company_id,company_name,isEnable,address,suburb_district,ward,postcode,state_province,country,description,policy,condition_to_book,logo_path,created_by,creation_date,last_updated_by,last_update_date from ocs.companies where company_id = ?", id)
	if err != nil {
		log.Println("companyMdl.find.go: All() err = ", err)
	}
	row := Company{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}

	rs.Scan(&row.CompanyId, &row.CompanyName, &row.IsEnable, &row.Address, &row.SuburbDistrict, &row.Ward, &row.Postcode, &row.StateProvince, &row.Country, &row.Description, &row.Policy, &row.ConditionToBook, &row.LogoPath, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time

	return row, err
}
