package skinRequestCtrl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/disintegration/imaging"

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

	img, err := jpeg.Decode(part)
	utils.LogError(" ==========> decode to img ", err)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(" \n\n\n\n===========> img = ", img, " err = ", err)
	dstImage800 := imaging.Resize(img, 200, 0, imaging.Lanczos)
	log.Println(" ======> dstImage800 = ", dstImage800, " err = ", err)
	imaging.Save(dstImage800, "./smallphotosupload/"+part.FileName())

	goGroup.Done()
	return nil
}

func saveToFile2(part *multipart.Part) error {

	start := time.Now()
	dst, err := os.Create("./photosupload/" + part.FileName())
	defer dst.Close()
	utils.LogError(" Create file ", err)
	if err != nil {
		return err
	}

	if _, err := io.Copy(dst, part); err != nil {
		utils.LogError(" Copy file ", err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	log.Printf("duration to create file %s = %s", part.FileName(), time.Since(start))

	src, err := imaging.Open("./photosupload/" + part.FileName())
	utils.LogError(" ==========> decode to img ", err)
	dstImage800 := imaging.Resize(src, 200, 0, imaging.Lanczos)
	imaging.Save(dstImage800, "./smallphotosupload/"+part.FileName())

	return nil
}

func saveToDB(goGroup *sync.WaitGroup, moleReq patientAppointmentMdl.MoleRequest) error {
	_, err := moleReq.SubmitMoles()
	utils.LogError("Submit mole request ", err)
	goGroup.Done()
	return nil
}

func SubmitMoles(w http.ResponseWriter, r *http.Request) {
	log := logger.Log
	start := time.Now()
	//wg := new(sync.WaitGroup)
	moles := patientAppointmentMdl.MoleRequest{}
	reader, err := r.MultipartReader()

	utils.LogError(" ParseMultipartForm ", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy each part to destination.
	//parts := []*multipart.Part{}
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
		//parts = append(parts, part)
		saveToFile2(part)
	}

	appt, err := moles.SubmitMoles()

	// wg.Add(1)
	// go saveToDB(wg, moles)

	// for _, part := range parts {
	// 	wg.Add(1)
	// 	go saveToFile(wg, part)
	// }

	//wg.Wait()
	log.Infof("total duration = %s", time.Since(start))

	utils.APIResponse(w, err, appt)

}
