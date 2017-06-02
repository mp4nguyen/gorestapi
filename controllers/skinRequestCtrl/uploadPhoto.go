package skinRequestCtrl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"

	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/patientAppointmentMdl"
	"bitbucket.org/restapi/utils"
)

func saveToFile(goGroup *sync.WaitGroup, part *multipart.Part) error {

	start := time.Now()
	dst, err := os.Create("./photosupload/" + part.FileName())
	defer dst.Close()
	utils.LogError(" Create file ", err)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	if _, err := io.Copy(dst, part); err != nil {
		utils.LogError(" Copy file ", err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	log.Printf("duration to create file %s = %s", part.FileName(), time.Since(start))
	goGroup.Done()
	return nil
}

func saveToDB(goGroup *sync.WaitGroup, moleReq patientAppointmentMdl.MoleRequest) error {
	err := moleReq.SubmitMoles()
	utils.LogError("Submit mole request ", err)
	goGroup.Done()
	return nil
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	start := time.Now()
	wg := new(sync.WaitGroup)
	moles := patientAppointmentMdl.MoleRequest{}
	reader, err := r.MultipartReader()

	utils.LogError(" ParseMultipartForm ", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy each part to destination.
	parts := []*multipart.Part{}
	for {
		part, err := reader.NextPart()
		//utils.LogError(" ParseMultipartForm ", err)
		if err == io.EOF {
			break
		}
		fmt.Println("part = ", part)
		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			buf := new(bytes.Buffer)
			buf.ReadFrom(part)
			s := buf.String()
			fmt.Println(" not file  = ", part.FormName(), "  buffer = ", s)
			json.Unmarshal(buf.Bytes(), &moles)
			fmt.Println("mole obj = ", moles)

			continue
		}
		parts = append(parts, part)
		//saveToFile2(part)
	}

	wg.Add(1)
	go saveToDB(wg, moles)

	for _, part := range parts {
		wg.Add(1)
		go saveToFile(wg, part)
	}
	//log.Print(" =======> parts = ", parts)
	//display success message.

	wg.Wait()
	log.Infof("total duration = %s", time.Since(start))

	fmt.Fprintf(w, "Upload successful.")

}
