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

func Signup2(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	member := accountMdl.PatientProfile{}

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
		creatingPerson.Title = member.Title
		creatingPerson.FirstName = member.FirstName
		creatingPerson.LastName = member.LastName
		creatingPerson.Dob, err = time.Parse(time.RFC3339, member.Dob)
		utils.LogError("parse dob time err", err)
		creatingPerson.Gender = member.Gender
		creatingPerson.Mobile = member.Mobile
		creatingPerson.Address = member.Address
		creatingPerson.SuburbDistrict = member.SuburbDistrict
		creatingPerson.StateProvince = member.StateProvince
		creatingPerson.Country = member.Country
		creatingPerson.Postcode = member.Postcode
		creatingPerson.Email = member.Email
		creatingPerson.GPFirstName = member.GPFirstName
		creatingPerson.GPLastName = member.GPLastName
		creatingPerson.GPContact = member.GPContactNumber
		creatingPerson.ClinicName = member.Clinic
		creatingPerson.MedicareNo = member.MedicareNo
		creatingPerson.MedicareRef = member.MedicareRef
		creatingPerson.MedicareExpired, err = time.Parse(time.RFC3339, member.MedicareExpired)
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
		creatingAccount.Username = member.Username
		creatingAccount.Password, err = utils.HashPassword(member.Password)
		creatingAccount.Email = member.Email
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
	login.Username = member.Username
	login.Password = member.Password
	_, acc, _ := login.CheckAccount()
	memberRes.Account = acc
	memberRes.IsSuccess = true
	output, err = json.Marshal(memberRes)
	utils.LogError("failed to json marshal memberRes", err)
	fmt.Fprintln(w, string(output))

}
