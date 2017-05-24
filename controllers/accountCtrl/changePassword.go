package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accountMdl"
	"bitbucket.org/restapi/models/patientMdl"
	"bitbucket.org/restapi/models/personMdl"
	"bitbucket.org/restapi/utils"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
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

	tx, err := db.GetDB().Begin()
	utils.ErrorHandler("Error when creating transaction", err, nil)
	if err == nil {
		creatingPerson := personMdl.Person{}
		creatingPerson.Title = member.Baseinfo.Title
		creatingPerson.FirstName = member.Baseinfo.FirstName
		creatingPerson.LastName = member.Baseinfo.LastName
		creatingPerson.Dob, err = time.Parse(time.RFC3339, member.Baseinfo.Dob)
		utils.ErrorHandler("parse time err", err, tx)
		creatingPerson.Gender = member.Baseinfo.Gender
		creatingPerson.Occupation = member.Baseinfo.Occupation
		creatingPerson.Mobile = member.Contact.Phone
		creatingPerson.Address = member.Contact.Address
		creatingPerson.SuburbDistrict = member.Contact.Suburb
		creatingPerson.StateProvince = member.Contact.State
		creatingPerson.Country = member.Contact.Country
		creatingPerson.Postcode = member.Contact.Postcode
		creatingPerson.Email = member.Signup.Email
		creatingPerson.IsEnable = 1
		creatingPerson.IsPatient = 1
		output, _ = json.Marshal(creatingPerson)
		log.Infof(" creating Person object  = %s", string(output))

		noOfPerson, lastPersonId, creatingPersonErr := creatingPerson.Create(tx)
		utils.ErrorHandler("Failed to create person", creatingPersonErr, tx)
		log.Infof("noOfPerson=%s, lastPersonId=%s, creatingPersonErr=%s", noOfPerson, lastPersonId, creatingPersonErr)

		creatingPatient := patientMdl.Patient{}
		creatingPatient.IsEnable = 1
		creatingPatient.PersonId = int(lastPersonId)
		creatingPatient.UserId = 1
		noOfPatient, lastPatientId, creatingPatientErr := creatingPatient.Create(tx)
		utils.ErrorHandler("Failed to create patient", creatingPatientErr, tx)
		log.Infof("noOfPatient=%s, lastPatientId=%s, creatingPatientErr=%s", noOfPatient, lastPatientId, creatingPatientErr)

		creatingAccount := accountMdl.Account{}
		creatingAccount.Username = member.Signup.Username
		creatingAccount.Password, err = utils.HashPassword(member.Signup.Password)
		creatingAccount.Email = member.Signup.Email
		creatingAccount.PersonId = int(lastPersonId)
		creatingAccount.PatientId = int(lastPatientId)
		creatingAccount.IsEnable = 1
		creatingAccount.UserType = "PATIENT"
		noOfAcc, lastAccId, creatingAccErr := creatingAccount.Create(tx)
		utils.ErrorHandler("Failed to create ACCOUNT", creatingAccErr, tx)
		log.Infof("noOfAcc=%s, lastAccId=%s, creatingAccErr=%s", noOfAcc, lastAccId, creatingAccErr)
		tx.Commit()
	}

	fmt.Fprintln(w, "Doing...")

}
