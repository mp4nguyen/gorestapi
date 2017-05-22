package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/models/accountMdl"
)

func LoginAT(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
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
	if err != nil {
		log.Errorf("Something went wrong (%s)!", err)
	}

	isMatch, acc, errCheckAccount := login.CheckAccount()

	if isMatch {
		acc.FetchPersonForAccount()
		output, err := json.Marshal(acc)
		log.Infof("Infor from client = %s", string(output))
		if err != nil {
			log.Errorf("Something went wrong (%s)!", err)
		}

		at, _ := accessTokenMdl.Create(1)
		log.Infof(" at = %s", at)
		fmt.Fprintln(w, string(output))
	} else {
		fmt.Fprintln(w, errCheckAccount)
	}

}
