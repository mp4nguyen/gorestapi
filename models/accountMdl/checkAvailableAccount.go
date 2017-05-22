package accountMdl

import (
	"strings"

	"bitbucket.org/restapi/db"
)

func (m Login) CheckAvailableAccount() (checkAvailableRes CheckAvailableRes) {

	sqlString := "select email,id,username from ocs.accounts where username=? or email=?"

	//log.Println(" m.Username = ", m.Username, " m.email = ", m.Email)
	rs := db.GetDB().QueryRow(sqlString, m.Username, m.Email)

	row := Account{}
	res := CheckAvailableRes{}

	rs.Scan(&row.Email, &row.Id, &row.Username)
	//log.Println("row = ", row)
	if strings.ToLower(row.Email) == strings.ToLower(m.Email) {
		res.IsAvailable = false
		res.Reason = m.Email + " has existed in the system"
		return res
	} else if strings.ToLower(row.Username) == strings.ToLower(m.Username) {
		res.IsAvailable = false
		res.Reason = m.Username + " has existed in the system"
		return res
	} else {
		res.IsAvailable = true
		return res
	}

}
