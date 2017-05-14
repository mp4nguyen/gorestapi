package accountCtrl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/restapi/cli/accountMdl"
)

func Find(w http.ResponseWriter, r *http.Request) {
	data, err := accountMdl.Find()
	if err != nil {
		fmt.Println(err)
	}
	output, _ := json.Marshal(data)
	fmt.Fprintln(w, string(output))
}
