package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"bitbucket.org/restapi/models/accountMdl"
	"bitbucket.org/restapi/models/companyMdl"
)

func Find(w http.ResponseWriter, r *http.Request) {

	calParams := accountMdl.Account{}
	dec := json.NewDecoder(r.Body)
	log.Println("dec Body = ", dec)
	//fmt.Println(dec)
	//fmt.Println(r.FormValue("id"))
	for {
		if err := dec.Decode(&calParams); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	outputJsonParam, errJsonParam := json.Marshal(calParams)
	log.Println(string(outputJsonParam))
	if errJsonParam != nil {
		fmt.Println("Something went wrong! errJsonParam = ", errJsonParam)
	}

	data, err := accountMdl.Find("", "")
	if err != nil {
		fmt.Println(err)
	}

	data.FetchPerson()

	companies, err := companyMdl.Find("", "")
	companies.FetchClinic()
	output, _ := json.Marshal(companies)
	fmt.Fprintln(w, string(output))

	// ////////test create///
	// accounts := accountMdl.Accounts{}
	// accounts = append(accounts, accountMdl.Account{UserType: "PATIENT", IsEnable: 1, Id: 0, Username: "testtesttest1", Password: "Passssssss"})
	// accounts = append(accounts, accountMdl.Account{UserType: "PATIENT", IsEnable: 1, Id: 0, Username: "testtesttest2", Password: "Passssssss"})
	// noOfRows, lastId, _ := accountMdl.Create(accounts)
	// fmt.Println("created ", noOfRows)
	// acc, _ := accountMdl.FindById(lastId)
	//
	// outputJsonParam, errJsonParam = json.Marshal(acc)
	// log.Println("Just created row = ", string(outputJsonParam))
	// if errJsonParam != nil {
	// 	fmt.Println("Something went wrong! errJsonParam = ", errJsonParam)
	// }

	//////////////////////

}
