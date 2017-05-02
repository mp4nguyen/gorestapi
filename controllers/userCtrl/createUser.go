package userCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"bitbucket.org/restapi/models/userMdl"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	NewUser := userMdl.User{}

	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&NewUser); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	output, err := json.Marshal(NewUser)
	log.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	q, err := userMdl.Create(NewUser)
	fmt.Println(q)
}
