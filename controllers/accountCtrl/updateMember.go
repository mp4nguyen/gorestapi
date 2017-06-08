package accountCtrl

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accountMdl"
	"bitbucket.org/restapi/models/personMdl"
	"bitbucket.org/restapi/utils"
)

func UpdateMember(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	member := accountMdl.Member{}

	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&member); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	output, err := json.Marshal(member)
	log.Infof("Infor from client = %s", string(output))
	utils.ErrorHandler("Json.Marshal for req body", err, nil)

	updatingMember := personMdl.UpdatingMember{}
	updatingMember.PersonId = member.Baseinfo.PersonId
	updatingMember.Title = member.Baseinfo.Title
	updatingMember.FirstName = member.Baseinfo.FirstName
	updatingMember.LastName = member.Baseinfo.LastName
	updatingMember.Occupation = member.Baseinfo.Occupation
	updatingMember.Gender = member.Baseinfo.Gender
	updatingMember.Dob, err = time.Parse(time.RFC3339, member.Baseinfo.Dob)
	utils.LogError("Parse DOB ", err)
	updatingMember.Address = member.Contact.Address
	updatingMember.SuburbDistrict = member.Contact.Suburb
	updatingMember.StateProvince = member.Contact.State
	updatingMember.Postcode = member.Contact.Postcode
	updatingMember.Country = member.Contact.Country
	updatingMember.Mobile = member.Contact.Phone
	updatingMember.GPFirstName = member.GP.FirstName
	updatingMember.GPLastName = member.GP.LastName
	updatingMember.ClinicName = member.GP.Clinic
	updatingMember.GPContact = member.GP.ContactNumber
	updatingMember.MedicareNo = member.GP.MedicareNo
	updatingMember.MedicareRef = member.GP.MedicareRef
	updatingMember.MedicareExpired, err = time.Parse(time.RFC3339, member.GP.MedicareExpired)
	utils.LogError("Parse MedicareExpired ", err)
	noOfRow, err := db.Update("people", updatingMember, nil)
	utils.LogError("Update member", err)
	log.Infof("Updated people %s row effected", noOfRow)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprintln(w, err)
	// } else {
	// 	output, err := json.Marshal(updatingMember)
	// 	utils.ErrorHandler("Json.Marshal for req body", err, nil)
	// 	fmt.Fprintln(w, string(output))
	// }

	utils.APIResponse(w, err, updatingMember)
}
