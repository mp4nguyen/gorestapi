package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"bitbucket.org/restapi/models/accountMdl"
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

	data, err := accountMdl.Find()
	if err != nil {
		fmt.Println(err)
	}
	output, _ := json.Marshal(data)
	fmt.Fprintln(w, string(output))
}
