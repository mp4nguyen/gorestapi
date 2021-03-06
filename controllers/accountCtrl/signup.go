package accountCtrl

import (
	"database/sql"
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

func Signup(w http.ResponseWriter, r *http.Request) {
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

	memberRes := personMdl.CreatedMemberRes{}

	db.Transaction(func(tx *sql.Tx) error {
		creatingPerson := personMdl.Person{}
		creatingPerson.Title = member.Baseinfo.Title
		creatingPerson.FirstName = member.Baseinfo.FirstName
		creatingPerson.LastName = member.Baseinfo.LastName
		creatingPerson.Dob, err = time.Parse(time.RFC3339, member.Baseinfo.Dob)
		utils.LogError("parse dob time err", err)
		creatingPerson.Gender = member.Baseinfo.Gender
		creatingPerson.Occupation = member.Baseinfo.Occupation
		creatingPerson.Mobile = member.Contact.Phone
		creatingPerson.Address = member.Contact.Address
		creatingPerson.SuburbDistrict = member.Contact.Suburb
		creatingPerson.StateProvince = member.Contact.State
		creatingPerson.Country = member.Contact.Country
		creatingPerson.Postcode = member.Contact.Postcode
		creatingPerson.Email = member.Signup.Email
		creatingPerson.GPFirstName = member.GP.FirstName
		creatingPerson.GPLastName = member.GP.LastName
		creatingPerson.GPContact = member.GP.ContactNumber
		creatingPerson.ClinicName = member.GP.Clinic
		creatingPerson.MedicareNo = member.GP.MedicareNo
		creatingPerson.MedicareRef = member.GP.MedicareRef
		creatingPerson.MedicareExpired, err = time.Parse(time.RFC3339, member.GP.MedicareExpired)
		utils.LogError("parse MedicareExpired time err", err)
		err = nil
		creatingPerson.IsEnable = 1
		creatingPerson.IsPatient = 1
		output, _ = json.Marshal(creatingPerson)
		log.Infof(" creating Person object  = %s", string(output))

		noOfPerson, lastPersonId, creatingPersonErr := creatingPerson.Create(tx)
		utils.LogError("Failed to create person", creatingPersonErr)
		if creatingPersonErr != nil {
			return creatingPersonErr
		}
		log.Infof("noOfPerson=%s, lastPersonId=%s, creatingPersonErr=%s", noOfPerson, lastPersonId, creatingPersonErr)

		creatingPatient := patientMdl.Patient{}
		creatingPatient.IsEnable = 1
		creatingPatient.PersonId = int(lastPersonId)
		creatingPatient.UserId = 1
		noOfPatient, lastPatientId, creatingPatientErr := creatingPatient.Create(tx)
		utils.LogError("Failed to create patient", creatingPatientErr)
		if creatingPatientErr != nil {
			return creatingPatientErr
		}
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
		utils.LogError("Failed to create ACCOUNT", creatingAccErr)
		log.Infof("noOfAcc=%s, lastAccId=%s, creatingAccErr=%s", noOfAcc, lastAccId, creatingAccErr)

		memberRes.PatientId = int(lastPatientId)
		memberRes.PersonId = int(lastPersonId)

		return creatingAccErr
	})

	// tx, err := db.GetDB().Begin()
	// utils.ErrorHandler("Error when creating transaction", err, nil)
	// if err == nil {
	//
	// 	tx.Commit()
	// }
	if err != nil {
		memberRes.IsSuccess = false
		memberRes.Reason = err.Error()
	}

	login := accountMdl.Login{}
	login.Username = member.Signup.Username
	login.Password = member.Signup.Password
	_, acc, _ := login.CheckAccount()
	memberRes.Account = acc
	memberRes.IsSuccess = true
	output, err = json.Marshal(memberRes)
	utils.LogError("failed to json marshal memberRes", err)
	fmt.Fprintln(w, string(output))

}
