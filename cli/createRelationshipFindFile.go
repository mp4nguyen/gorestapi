package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createRelationshipFindFile(c *ishell.Context, folderName string, modelName string, detailModelName string, masterColumnKey string, detailColumnKey string, masterTableForeignKey string, detailTableForeignKey string, relationshipType string, relationShipName string) {
	////create find.go////
	findFile := []byte("")
	appendToBytes(&findFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))
	appendToBytes(&findFile, fmt.Sprintf("import \"log\"\n"))
	appendToBytes(&findFile, fmt.Sprintf("import \"bitbucket.org/restapi/db\"\n\n"))

	///create fetch data for single data

	appendToBytes(&findFile, fmt.Sprintf("func (m *%s)Fetch%sFor%s()(err error){\n", modelName, detailModelName, modelName))

	if relationshipType == "1" {
		appendToBytes(&findFile, fmt.Sprintf("\twhereCondition := \"%s = strconv.Itoa(m.%s)\"\n", detailColumnKey, fieldName(detailColumnKey)))
	} else if relationshipType == "2" {
		appendToBytes(&findFile, fmt.Sprintf("\twhereCondition := \"%s = strconv.Itoa(m.%s)\"\n", detailTableForeignKey, fieldName(detailTableForeignKey)))
	}

	if relationshipType == "1" {
		appendToBytes(&findFile, fmt.Sprintf("\ttempMapData, err := %sMdl.MapFind(\"%s\",whereCondition, \"%s\")\n", LcFirst(detailModelName), fieldName(detailColumnKey), detailColumnKey))
		appendToBytes(&findFile, fmt.Sprintf("\t\ttempData, ok := tempMapData[strconv.Itoa(m.%s)]\n", fieldName(detailColumnKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\tif ok {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\tif len(tempData) > 0 {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\tm.%s = tempData[0]\n", relationShipName))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\t}\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t}\n"))
	} else if relationshipType == "2" {
		appendToBytes(&findFile, fmt.Sprintf("\ttempMapData, err := %sMdl.MapFind(%s,whereCondition, \"%s\")\n", LcFirst(detailModelName), fieldName(detailTableForeignKey), detailTableForeignKey))
		appendToBytes(&findFile, fmt.Sprintf("\t\ttempData, ok := tempMapData[strconv.Itoa(m.%s)]\n", fieldName(detailTableForeignKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\tif ok {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\tif len(tempData) > 0 {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\tm.%s = tempData[0]\n", relationShipName))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\t}\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t}\n"))
	}

	appendToBytes(&findFile, fmt.Sprintf("\treturn err\n"))

	appendToBytes(&findFile, "}\n")

	///create fetch data for array of data
	appendToBytes(&findFile, fmt.Sprintf("func (m *%ss)Fetch%sFor%ss()(err error){\n", modelName, detailModelName, modelName))

	if relationshipType == "1" {
		appendToBytes(&findFile, fmt.Sprintf("\tforeignKeys := map[string]string{}\n"))
		appendToBytes(&findFile, fmt.Sprintf("\twhereCondition := \"%s in (\"\n", detailColumnKey))
	} else if relationshipType == "2" {
		appendToBytes(&findFile, fmt.Sprintf("\twhereCondition := \"%s in (\"\n", detailTableForeignKey))
	}

	appendToBytes(&findFile, fmt.Sprintf("\tfor _, row := range *m {\n"))

	if relationshipType == "1" {
		appendToBytes(&findFile, fmt.Sprintf("\t\t_, ok := foreignKeys[strconv.Itoa(row.%s)]\n", fieldName(masterTableForeignKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\tif !ok {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\tforeignKeys[strconv.Itoa(row.%s)] = strconv.Itoa(row.%s)\n", fieldName(masterTableForeignKey), fieldName(masterTableForeignKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\twhereCondition = whereCondition + strconv.Itoa(row.%s) + \",\"\n", fieldName(masterTableForeignKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\t}\n"))
	} else if relationshipType == "2" {
		appendToBytes(&findFile, fmt.Sprintf("\t\t\twhereCondition = whereCondition + strconv.Itoa(row.%s) + \",\"\n", fieldName(masterColumnKey)))
	}

	appendToBytes(&findFile, fmt.Sprintf("\t}\n"))

	appendToBytes(&findFile, fmt.Sprintf("\twhereCondition = whereCondition[0:len(whereCondition)-1] + \")\"\n"))

	if relationshipType == "1" {
		appendToBytes(&findFile, fmt.Sprintf("\ttempMapData, err := %sMdl.MapFind(\"%s\",whereCondition, \"%s\")\n", LcFirst(detailModelName), fieldName(detailColumnKey), detailColumnKey))
		appendToBytes(&findFile, fmt.Sprintf("\tfor _, row := range *m {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\ttempData, ok := tempMapData[strconv.Itoa(row.%s)]\n", fieldName(detailColumnKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\tif ok {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\tif len(tempData) > 0 {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\trow.%s = tempData[0]\n", relationShipName))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\t}\n"))

		appendToBytes(&findFile, fmt.Sprintf("\t\t}\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t}\n"))
	} else if relationshipType == "2" {
		appendToBytes(&findFile, fmt.Sprintf("\ttempMapData, err := %sMdl.MapFind(%s,whereCondition, \"%s\")\n", LcFirst(detailModelName), fieldName(detailTableForeignKey), detailTableForeignKey))
		appendToBytes(&findFile, fmt.Sprintf("\tfor _, row := range *m {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\ttempData, ok := tempMapData[strconv.Itoa(row.%s)]\n", fieldName(detailTableForeignKey)))
		appendToBytes(&findFile, fmt.Sprintf("\t\tif ok {\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t\t\trow.%s = tempData\n", relationShipName))
		appendToBytes(&findFile, fmt.Sprintf("\t\t}\n"))
		appendToBytes(&findFile, fmt.Sprintf("\t}\n"))
	}

	appendToBytes(&findFile, fmt.Sprintf("\treturn err\n"))

	appendToBytes(&findFile, "}\n")

	errFindFile := ioutil.WriteFile(fmt.Sprintf("outputs/%s/fetch%s.go", folderName, detailModelName), findFile, 0644)
	if errFindFile != nil {
		c.Println("Error while writing to file err = ", errFindFile)
	}

}
