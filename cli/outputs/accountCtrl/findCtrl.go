package accountCtrl

import "time"
import "log"
import "bitbucket.org/restapi/db"

func Find(w http.ResponseWriter, r *http.Request) {
	data, err := accountMdl.Find()
	if err != nil {
		fmt.Println(err)
	}
	output, _ := json.Marshal(data)
	fmt.Fprintln(w, string(output))
}
