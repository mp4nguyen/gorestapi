package userCtrl

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AfterLogin(w http.ResponseWriter, r *http.Request) {

	//r.Context().Value("UserId")
	//r.Context().Value("UserId")
	//fmt.Fprintln(w, "string(json) UserId = ")
	headOrTails := rand.Intn(2)

	if headOrTails == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "Go! slow %v", headOrTails)
		fmt.Printf("Go! slow %v", headOrTails)
		return
	}

	fmt.Fprintf(w, "Go! quick %v", headOrTails)
	fmt.Printf("Go! quick %v", headOrTails)
	return

}
