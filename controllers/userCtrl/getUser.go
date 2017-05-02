package userCtrl

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/restapi/models/userMdl"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Pragma", "no-cache")

	urlParams := mux.Vars(r)
	id := urlParams["id"]
	userId, err := strconv.ParseInt(id, 10, 64)

	data, err := userMdl.One(userId)
	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "No such user")
	case err != nil:
		log.Fatal(err)
	default:
		output, _ := json.Marshal(data)
		fmt.Fprintf(w, string(output))
	}
}
