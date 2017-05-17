package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createFindByIdFile(c *ishell.Context, folderName string, schemaName string, tableName string, modelName string, outputFields string, queryFields string, columnKey string, declareDateFields string, assignDateFields string) {
	////create find.go////
	findFile := []byte("")
	appendToBytes(&findFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))
	appendToBytes(&findFile, fmt.Sprintf("import \"log\"\n"))
	appendToBytes(&findFile, fmt.Sprintf("import \"bitbucket.org/restapi/db\"\n\n"))

	appendToBytes(&findFile, fmt.Sprintf("func FindById(id int64)(%ss %s,err error){\n", LcFirst(modelName), modelName))

	appendToBytes(&findFile, fmt.Sprintf("\trs := db.GetDB().QueryRow(\"select %s from %s.%s where %s = ?\",id)\n", queryFields, schemaName, tableName, columnKey))
	appendToBytes(&findFile, "\tif err != nil {\n")
	appendToBytes(&findFile, fmt.Sprintf("\t\tlog.Println(\"%sMdl.find.go: All() err = \", err)\n", LcFirst(modelName)))
	appendToBytes(&findFile, "\t}\n")

	appendToBytes(&findFile, fmt.Sprintf("\trow := %s{}\n", modelName))
	appendToBytes(&findFile, fmt.Sprintf("\t\t%s\n", declareDateFields))
	appendToBytes(&findFile, fmt.Sprintf("\trs.Scan(%s)\n", outputFields))
	appendToBytes(&findFile, fmt.Sprintf("\t\t%s\n", assignDateFields))

	appendToBytes(&findFile, "\treturn row, err\n")
	appendToBytes(&findFile, "}\n")

	errFindFile := ioutil.WriteFile(folderName+"/findById.go", findFile, 0644)
	if errFindFile != nil {
		c.Println("Error while writing to file err = ", errFindFile)
	}

}
