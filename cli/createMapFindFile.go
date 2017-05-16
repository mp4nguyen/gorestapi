package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createMapFindFile(c *ishell.Context, folderName string, schemaName string, tableName string, modelName string, outputFields string, queryFields string, declareDateFields string, assignDateFields string, keyField string) {
	////create find.go////
	findFile := []byte("")
	appendToBytes(&findFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))

	appendToBytes(&findFile, fmt.Sprintf("import \"log\"\n"))
	appendToBytes(&findFile, fmt.Sprintf("import \"bitbucket.org/restapi/db\"\n\n"))

	appendToBytes(&findFile, fmt.Sprintf("func MapFind(where string, orderBy string)(%ss map[string]%s,err error){\n", LcFirst(modelName), modelName))

	appendToBytes(&findFile, fmt.Sprintf("\tsqlString := \"select %s from %s.%s\"\n", queryFields, schemaName, tableName))

	appendToBytes(&findFile, "\tif len(where) > 0 {\n")
	appendToBytes(&findFile, "\t\tsqlString += (\" where \" + where)\n")
	appendToBytes(&findFile, "\t}\n")

	appendToBytes(&findFile, "\tif len(orderBy) > 0 {\n")
	appendToBytes(&findFile, "\t\tsqlString += (\" order by \" + orderBy)\n")
	appendToBytes(&findFile, "\t}\n")

	appendToBytes(&findFile, fmt.Sprintf("\trows, err := db.GetDB().Query(sqlString)\n"))

	appendToBytes(&findFile, "\tif err != nil {\n")
	appendToBytes(&findFile, fmt.Sprintf("\t\tlog.Println(\"%sMdl.find.go: All() err = \", err)\n", LcFirst(modelName)))
	appendToBytes(&findFile, "\t}\n")
	appendToBytes(&findFile, "\tdefer rows.Close()\n\n")

	appendToBytes(&findFile, fmt.Sprintf("\tresponse := map[string]%s{}\n", modelName))
	appendToBytes(&findFile, "\tfor rows.Next() {\n")
	appendToBytes(&findFile, fmt.Sprintf("\t\trow := %s{}\n", modelName))

	appendToBytes(&findFile, fmt.Sprintf("\t\t%s\n", declareDateFields))
	appendToBytes(&findFile, fmt.Sprintf("\t\trows.Scan(%s)\n", outputFields))
	appendToBytes(&findFile, fmt.Sprintf("\t\t%s\n", assignDateFields))
	appendToBytes(&findFile, fmt.Sprintf("\t\tresponse[strconv.Itoa(row.%s)] = row\n", keyField))
	appendToBytes(&findFile, "\t}\n\n")
	appendToBytes(&findFile, "\treturn response, err\n")
	appendToBytes(&findFile, "}\n")

	errFindFile := ioutil.WriteFile(folderName+"/mapFind.go", findFile, 0644)
	if errFindFile != nil {
		c.Println("Error while writing to file err = ", errFindFile)
	}

}
