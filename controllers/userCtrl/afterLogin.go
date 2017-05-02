package userCtrl

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func AfterLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "string(json)")
}
