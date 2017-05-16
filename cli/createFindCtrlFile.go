package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createFindCtrlFile(c *ishell.Context, folderCtrlName string, modelName string) {
	findCtrlFile := []byte("")
	appendToBytes(&findCtrlFile, fmt.Sprintf("package %sCtrl\n\n", LcFirst(modelName)))
	appendToBytes(&findCtrlFile, fmt.Sprintf("import \"time\"\n"))
	appendToBytes(&findCtrlFile, fmt.Sprintf("import \"log\"\n"))
	appendToBytes(&findCtrlFile, fmt.Sprintf("import \"bitbucket.org/restapi/db\"\n\n"))

	appendToBytes(&findCtrlFile, "func Find(w http.ResponseWriter, r *http.Request) {\n")
	appendToBytes(&findCtrlFile, fmt.Sprintf("\tdata, err := %sMdl.Find()\n", LcFirst(modelName)))
	appendToBytes(&findCtrlFile, "\tif err != nil {\n")
	appendToBytes(&findCtrlFile, "\t\tfmt.Println(err)\n")
	appendToBytes(&findCtrlFile, "\t}\n")
	appendToBytes(&findCtrlFile, "\toutput, _ := json.Marshal(data)\n")
	appendToBytes(&findCtrlFile, "\tfmt.Fprintln(w, string(output))\n")
	appendToBytes(&findCtrlFile, "}\n")

	errFindCtrlFile := ioutil.WriteFile(folderCtrlName+"/findCtrl.go", findCtrlFile, 0644)
	if errFindCtrlFile != nil {
		c.Println("Error while writing to file err = ", errFindCtrlFile)
	}
}
