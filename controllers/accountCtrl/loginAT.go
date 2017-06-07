package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"bitbucket.org/restapi/logger"
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
		output, err := json.Marshal(acc)
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
