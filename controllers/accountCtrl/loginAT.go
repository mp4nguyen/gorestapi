package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/models/accountMdl"
	"bitbucket.org/restapi/models/personMdl"
	"bitbucket.org/restapi/utils"
)

func LoginAT(w http.ResponseWriter, r *http.Request) {
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

	isMatch, acc, errCheckAccount := login.CheckAccount()
	log.Infof("checked username and pass = %s", isMatch)

	if isMatch {
		acc.FetchPerson()
		at, err := accessTokenMdl.Create(acc.Id)
		utils.ErrorHandler("Accesstoken generated ", err, nil)

		acc.AccessToken = at
		log.Infof(" at = %s", at)
		/////prepare object to return to client
		patientAccRes := accountMdl.PatientAccountRes{
			PersonId:       acc.PersonId,
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
			Profiles:       []personMdl.Person{acc.Person},
			AccessToken:    acc.AccessToken,
		}

		loginRes.IsLogin = true
		loginRes.AccessToken = at
		loginRes.Account = patientAccRes
		output, err := json.Marshal(loginRes)
		log.Infof("Infor from client = %s", string(output))
		utils.ErrorHandler("json marshal ", err, nil)
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
