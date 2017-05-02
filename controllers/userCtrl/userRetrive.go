package userCtrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/restapi/models/userMdl"

	_ "github.com/go-sql-driver/mysql"
)

func UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("starting retrieval")
	start := 0
	limit := 10

	next := start + limit

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", "<http://localhost:8080/api/users?start="+string(next)+"; rel=\"next\"")

	data, err := userMdl.All()

	if err != nil {
		fmt.Println(err)
	}

	output, _ := json.Marshal(data)
	fmt.Fprintln(w, string(output))
}
