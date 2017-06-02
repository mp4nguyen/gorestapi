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

	"bitbucket.org/restapi/utils"
)

type moleRequest struct {
	LesionId             int      `json:"lesionId"`
	PersonId             int      `json:"personId"`
	Gender               string   `json:"gender"`
	Lesion               string   `json:"lesion"`
	IsFront              bool     `json:"isFront"`
	IsNew                bool     `json:"isNew"`
	IsGrowing            bool     `json:"isGrowing"`
	IsShapeOrChangeColor bool     `json:"isShapeOrChangeColor"`
	IsItchyOrBleeding    bool     `json:"isItchyOrBleeding"`
	IsTenderOrPainful    bool     `json:"isTenderOrPainful"`
	DoesItComeAndGo      bool     `json:"doesItComeAndGo"`
	Resource             []string `json:"resource"`
}

type moleRequests []moleRequest

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

func saveToFile2(part *multipart.Part) error {

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

	return nil
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// log := logger.Log
	// req := moleRequest{}
	//
	// dec := json.NewDecoder(r.Body)
	// for {
	//
	// 	if err := dec.Decode(&req); err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// output, err := json.Marshal(req)
	// log.Infof("Infor from client = %s", string(output))
	// utils.ErrorHandler("Json.Marshal for req body", err, nil)

	// //get the multipart reader for the request.

	start := time.Now()
	wg := new(sync.WaitGroup)
	moles := moleRequests{}
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

	for _, part := range parts {
		wg.Add(1)
		go saveToFile(wg, part)
	}
	//log.Print(" =======> parts = ", parts)
	//display success message.

	wg.Wait()
	log.Printf("total duration = %s", time.Since(start))

	fmt.Fprintf(w, "Upload successful.")

	////////////////////////////
	//fmt.Println("UploadHandler ................")
	// the FormFile function takes in the POST input id file

	//fmt.Println("UploadHandler ................ FormValue = ", r.FormValue("company"), " Form = ", r.Form)

	// errParse := r.ParseMultipartForm(0)
	// utils.LogError(" ParseMultipartForm ", errParse)
	// m := r.MultipartForm
	//
	// fmt.Println("UploadHandler ................m=", m)
	//
	// files := m.File["files"]
	// fmt.Println(1)
	// for i, _ := range files {
	// 	//for each fileheader, get a handle to the actual file
	// 	fmt.Println(" file i = ", i)
	// 	file, err := files[i].Open()
	// 	defer file.Close()
	// 	if err != nil {
	// 		fmt.Println(" open file err = ", err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	//create destination file making sure the path is writeable.
	// 	//f, err := os.OpenFile("./photosupload/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// 	dst, err := os.Create("./photosupload/" + files[i].Filename)
	// 	defer dst.Close()
	// 	if err != nil {
	// 		fmt.Println(" save file err = ", err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	//copy the uploaded file to the destination file
	// 	if _, err := io.Copy(dst, file); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	//
	// }
	//
	// fmt.Fprintf(w, "File uploaded successfully : ")

	/////////////////////////
	// fmt.Println("UploadHandler ................ FormValue = ", r.FormValue("company"), " Form = ", r.Form)
	//
	// file, header, err := r.FormFile("file")
	// utils.LogError(" FormFile ", err)
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// 	return
	// }
	//
	// defer file.Close()
	//
	// f, err := os.OpenFile("./photosupload/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// utils.LogError(" OpenFile ", err)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
	// 	return
	// }
	// defer f.Close()
	//
	// // write the content from POST to the file
	// _, err = io.Copy(f, file)
	// utils.LogError(" Copy ", err)
	// if err != nil {
	// 	fmt.Println("io.Copy ................", err)
	// 	fmt.Fprintln(w, err)
	// }
	//
	// fmt.Fprintf(w, "File uploaded successfully : ")
	// fmt.Fprintf(w, header.Filename)
}
