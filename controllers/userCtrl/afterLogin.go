package userCtrl

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func AfterLogin(w http.ResponseWriter, r *http.Request) {

	//r.Context().Value("UserId")
	//r.Context().Value("UserId")
	fmt.Fprintln(w, "string(json) UserId = ")
}
