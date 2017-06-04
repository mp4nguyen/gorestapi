package skinRequestCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/personMdl"
	"bitbucket.org/restapi/utils"
)

func GetAppointment(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	person := personMdl.Person{}

	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&person); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	person.FetchPatientAppointment()
	output, err := json.Marshal(person)
	log.Infof("Infor from client = %s", string(output))
	utils.ErrorHandler("json marshal ", err, nil)

	fmt.Fprintf(w, string(output))

}
