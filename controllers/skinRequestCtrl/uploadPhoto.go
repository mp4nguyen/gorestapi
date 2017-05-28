package skinRequestCtrl

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UploadHandler ................")
	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")
	fmt.Println("UploadHandler ................", file, header)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	f, err := os.OpenFile("./"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}
	defer f.Close()
	//io.Copy(f, file)

	// out, err := os.Create("/tmp/uploadedfile")
	// if err != nil {
	// 	fmt.Println("create:/tmp/uploadfile ................", err)
	// 	fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
	// 	return
	// }

	//defer f.Close()

	// write the content from POST to the file
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("io.Copy ................", err)
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)
}
