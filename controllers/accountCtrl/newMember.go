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
	"bitbucket.org/restapi/models/patientRelationshipMdl"
	"bitbucket.org/restapi/models/personMdl"
	"bitbucket.org/restapi/utils"
)

func NewMember(w http.ResponseWriter, r *http.Request) {
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

	err = db.Transaction(func(tx *sql.Tx) error {
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
		creatingPerson.GPFirstName = member.GP.FirstName
		creatingPerson.GPLastName = member.GP.LastName
		creatingPerson.GPContact = member.GP.ContactNumber
		creatingPerson.ClinicName = member.GP.Clinic
		creatingPerson.MedicareNo = member.GP.MedicareNo
		creatingPerson.MedicareRef = member.GP.MedicareRef
		creatingPerson.MedicareExpired, err = time.Parse(time.RFC3339, member.GP.MedicareExpired)
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

		creatingPatientRel := patientRelationshipMdl.PatientRelationship{}
		creatingPatientRel.FatherPersonId = member.Baseinfo.FatherPersonId
		creatingPatientRel.PatientId = int(lastPatientId)
		creatingPatientRel.PersonId = int(lastPersonId)
		creatingPatientRel.IsEnable = 1
		creatingPatientRel.RelationshipType = "FAMILY"
		noOfPatientRel, lastPatientRelId, creatingPatientRelErr := creatingPatientRel.Create(tx)
		utils.LogError("Failed to create patient relationship", creatingPatientRelErr)
		log.Infof("noOfPatient=%s, lastPatientId=%s, creatingPatientErr=%s", noOfPatientRel, lastPatientRelId, creatingPatientRelErr)

		memberRes.PatientId = int(lastPatientId)
		memberRes.PersonId = int(lastPersonId)
		return creatingPatientRelErr
	})

	if err != nil {
		memberRes.IsSuccess = false
		memberRes.Reason = err.Error()
	}

	output, err = json.Marshal(memberRes)
	utils.LogError("failed to json marshal memberRes", err)
	fmt.Fprintln(w, string(output))

}
