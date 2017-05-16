package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createModelFile(c *ishell.Context, folderName string, modelName string, modelFields string) {
	////create 0model.go////
	modelFile := []byte("")
	appendToBytes(&modelFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))
	appendToBytes(&modelFile, fmt.Sprintf("import \"time\"\n\n"))
	appendToBytes(&modelFile, fmt.Sprintf("type %s struct{\n", modelName))
	appendToBytes(&modelFile, modelFields)
	appendToBytes(&modelFile, fmt.Sprintf("\t}\n\n"))
	appendToBytes(&modelFile, fmt.Sprintf("type %ss []*%s", modelName, modelName))
	errModelFile := ioutil.WriteFile(folderName+"/0model.go", modelFile, 0644)
	if errModelFile != nil {
		c.Println("Error while writing to file err = ", errModelFile)
	}
}
