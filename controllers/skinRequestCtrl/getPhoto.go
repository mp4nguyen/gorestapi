package skinRequestCtrl

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/photoMdl"
	"bitbucket.org/restapi/utils"
)

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	photo := photoMdl.Photo{}

	dec := json.NewDecoder(r.Body)
	for {
		if err := dec.Decode(&photo); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	log.Info(" parameter = ", photo)

	// person.FetchPatientAppointment()

	// output, err := json.Marshal(photo)
	// log.Infof("Infor from client = %s", string(output))
	// utils.ErrorHandler("json marshal ", err, nil)
	log.Info("Will read file:", "./photosupload/"+photo.Uri)
	file, err := ioutil.ReadFile("./photosupload/" + photo.Uri)

	//log.Info(" file = ", file)
	utils.LogError(" loading file "+photo.Uri+" failed ", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(file)

	//w.Header().Set("Content-type", "image/jpg")
	//w.Write(file)

	fmt.Fprintf(w, imgBase64Str)

}
