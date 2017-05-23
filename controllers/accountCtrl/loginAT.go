package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/models/accountMdl"
	"bitbucket.org/restapi/utils"
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
	utils.ErrorHandler("json marshal ", err, nil)

	isMatch, acc, errCheckAccount := login.CheckAccount()
	log.Infof("checked username and pass = %s", isMatch)

	if isMatch {
		acc.FetchPerson()
		at, err := accessTokenMdl.Create(acc.Id)
		utils.ErrorHandler("Accesstoken generated ", err, nil)

		acc.AccessToken = at
		log.Infof(" at = %s", at)

		output, err := json.Marshal(acc)
		log.Infof("Infor from client = %s", string(output))
		utils.ErrorHandler("json marshal ", err, nil)

		fmt.Fprintln(w, string(output))
	} else {
		fmt.Fprintln(w, errCheckAccount)
	}

}
