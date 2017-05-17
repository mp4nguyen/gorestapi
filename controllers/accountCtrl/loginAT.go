package accountCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/models/accountMdl"
)

func LoginAT(w http.ResponseWriter, r *http.Request) {

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
	log.Println("Infor from client = ", string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	isMatch, acc, errCheckAccount := login.CheckAccount()

	fmt.Println("=====> ", isMatch, acc, errCheckAccount)

	if isMatch {
		acc.FetchPerson2()
		output, err := json.Marshal(acc)
		log.Println("Infor from client = ", string(output))
		if err != nil {
			fmt.Println("Something went wrong!")
		}

		at, _ := accessTokenMdl.Create(1)
		fmt.Println("at = ", at)
		fmt.Fprintln(w, string(output))
	} else {
		output, err := json.Marshal(errCheckAccount)
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		fmt.Fprintln(w, string(output))
	}

}
