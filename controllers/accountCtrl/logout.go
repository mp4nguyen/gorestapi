package accountCtrl

import (
	"fmt"
	"net/http"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/models/accessTokenMdl"

	_ "github.com/go-sql-driver/mysql"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	accessTokenHeader := r.Header["Accesstoken"]
	fmt.Println("JWT Middleware: accessTokenHeader = ", accessTokenHeader)

	if len(accessTokenHeader) > 0 {
		accessToken := r.Header["Accesstoken"][0]
		db.GetRedis().Del(accessToken)
		accessTokenMdl.ATCaching{}.DeleteAccessToken(accessToken)

	}

	fmt.Fprintln(w, "Logout successfully")
}
