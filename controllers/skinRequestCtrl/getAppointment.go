package skinRequestCtrl

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/patientAppointmentMdl"
	"bitbucket.org/restapi/utils"
)

type apptID struct {
	ApptID int `json:"apptId"`
}

func getPhoto(photoName string) string {
	log := logger.Log

	log.Info("Will read file:", "./photosupload/"+photoName)
	file, err := ioutil.ReadFile("./photosupload/" + photoName)

	//log.Info(" file = ", file)
	utils.LogError(" loading file "+photoName+" failed ", err)

	return base64.StdEncoding.EncodeToString(file)
}

func GetAppointment(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	appt := patientAppointmentMdl.PatientAppointment{}

	dec := json.NewDecoder(r.Body)
	for {
		if err := dec.Decode(&appt); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	appt.FetchRequest()
	appt.Requests.FetchPhoto()
	log.Info(" parameter = ", appt)
	// person.FetchPatientAppointment()
	res := patientAppointmentMdl.MoleRequestRes{}
	lesions := []patientAppointmentMdl.LesionRes{}
	//log.Info(" appt.PatientId = ", appt.PatientId)
	for _, req := range appt.Requests {
		//log.Info(" req.Data = ", req.Data)
		lesion := patientAppointmentMdl.LesionRes{}
		err := json.Unmarshal([]byte(req.Data), &lesion)
		utils.LogError("unmarshal json req data ", err)
		lesion.Lesion = getPhoto(lesion.Lesion)
		photostemp := []string{}
		for _, photo := range lesion.Photos {
			photostemp = append(photostemp, getPhoto(photo))
		}
		lesion.Resource = photostemp
		lesions = append(lesions, lesion)
		//log.Info(" lesion = ", lesion)
	}
	res.Lesions = lesions

	//log.Info("===> res = ", res)
	output, err := json.Marshal(res)
	//log.Infof("Infor from client = %s", string(output))
	utils.ErrorHandler("json marshal ", err, nil)

	fmt.Fprintf(w, string(output))

}
