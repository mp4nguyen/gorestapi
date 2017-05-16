package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createModelRelationshipFile(c *ishell.Context, folderName string, modelName string, relationType string, relationShipName string, detailModel string) {
	////create 0model.go////
	modelFile := []byte("")
	appendToBytes(&modelFile, "/*Please copy the property below to main model file: 0model.go to extend the relationship*/\n")
	appendToBytes(&modelFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))

	appendToBytes(&modelFile, fmt.Sprintf("type %s struct{\n", modelName))

	if relationType == "1" {
		appendToBytes(&modelFile, fmt.Sprintf("\t%s %sMdl.%s `json:\"%s\"`\n", relationShipName, LcFirst(detailModel), detailModel, LcFirst(relationShipName)))
	} else if relationType == "2" {
		appendToBytes(&modelFile, fmt.Sprintf("\t%s []%sMdl.%s `json:\"%ss\"`\n", relationShipName, LcFirst(detailModel), detailModel, LcFirst(relationShipName)))
	}

	appendToBytes(&modelFile, fmt.Sprintf("\t}\n\n"))

	errModelFile := ioutil.WriteFile(fmt.Sprintf("%s/0%s_%s.go", folderName, modelName, relationShipName), modelFile, 0644)
	if errModelFile != nil {
		c.Println("Error while writing to file err = ", errModelFile)
	}
}
