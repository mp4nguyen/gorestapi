package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accountMdl"
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
	if err != nil {
		log.Errorf("Something went wrong (%s)!", err)
	}

	fmt.Fprintln(w, "Doing...")

}
