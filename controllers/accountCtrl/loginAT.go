package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/models/accountMdl"
	"bitbucket.org/restapi/utils"
)

func LoginAT(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	log := logger.Log
	loginRes := accountMdl.LoginRes{}
	login := accountMdl.Login{}

	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&login); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	output, err := json.Marshal(login)
	log.Infof("Infor from client = %s", string(output))
	utils.ErrorHandler("json marshal ", err, nil)

	log.Infof(
		"duration0 = %s",
		time.Since(start),
	)
	isMatch, acc, errCheckAccount := login.CheckAccount()
	log.Infof("checked username and pass = %s", isMatch)

	log.Infof(
		"duration = %s",
		time.Since(start),
	)

	if isMatch {
		acc.FetchPerson()
		acc.Person.FetchPatientRelationshipV()
		at, err := accessTokenMdl.Create(acc.Id)
		utils.ErrorHandler("Accesstoken generated ", err, nil)

		log.Infof(
			"duration2 = %s",
			time.Since(start),
		)
		acc.AccessToken = at
		log.Infof(" at = %s", at)
		/////prepare object to return to client
		patientAccRes := accountMdl.PatientAccountRes{
			PersonId:       acc.PersonId,
			PatientId:      acc.PatientId,
			Username:       acc.Username,
			Title:          acc.Person.Title,
			FirstName:      acc.Person.FirstName,
			LastName:       acc.Person.LastName,
			Dob:            acc.Person.Dob,
			Gender:         acc.Person.Gender,
			Address:        acc.Person.Address,
			SuburbDistrict: acc.Person.SuburbDistrict,
			Ward:           acc.Person.Ward,
			Postcode:       acc.Person.Postcode,
			StateProvince:  acc.Person.StateProvince,
			Country:        acc.Person.Country,
			Profile:        acc.Person,
			AccessToken:    acc.AccessToken,
		}

		loginRes.IsLogin = true
		loginRes.AccessToken = at
		loginRes.Account = patientAccRes
		output, err := json.Marshal(loginRes)
		log.Infof("Infor from client = %s", string(output))
		utils.ErrorHandler("json marshal ", err, nil)

		log.Infof(
			"duration3 = %s",
			time.Since(start),
		)

		fmt.Fprintln(w, string(output))
	} else {
		loginRes.IsLogin = false
		loginRes.Reason = errCheckAccount.Error()
		output, err := json.Marshal(loginRes)
		log.Infof("Infor from client = %s", string(output))
		utils.ErrorHandler("json marshal ", err, nil)
		fmt.Fprintln(w, string(output))
	}

}
