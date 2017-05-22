package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accountMdl"
)

func CheckAvailableAccount(w http.ResponseWriter, r *http.Request) {

	log := logger.Log

	login := accountMdl.Login{}

	dec := json.NewDecoder(r.Body)
	for {
		if err := dec.Decode(&login); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Cannot parse Json body : ", err)
		}
	}
	output, err := json.Marshal(login)
	log.Infof("CheckAvailableAccount.go: Infor from client = %s", string(output))
	if err != nil {
		fmt.Println("Something went wrong!", err)
	}

	checkAvaiAccount := login.CheckAvailableAccount()

	output2, err2 := json.Marshal(checkAvaiAccount)
	if err2 != nil {
		fmt.Println("Something went wrong (json.Marshal)!", err2)
	}
	log.Infof("model res = %s", string(output2))
	fmt.Fprintln(w, string(output2))
}
