package main

import (
	"fmt"
	"io/ioutil"

	"github.com/abiosoft/ishell"
)

func createCreateFile(c *ishell.Context, folderName string, schemaName string, tableName string, modelName string, preFieldsForInsert string, modelFieldsForInsert string, queryFields string) {
	////create find.go////
	findFile := []byte("")
	appendToBytes(&findFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))

	appendToBytes(&findFile, fmt.Sprintf("import \"log\"\n"))
	appendToBytes(&findFile, fmt.Sprintf("import \"bitbucket.org/restapi/db\"\n\n"))

	appendToBytes(&findFile, fmt.Sprintf("func Create(inputs %ss) (noOfRows int64, lastId int64,err error) {\n", modelName))
	appendToBytes(&findFile, fmt.Sprintf("\tsqlStr := \"INSERT INTO %s(%s) VALUES \"\n", tableName, queryFields))
	appendToBytes(&findFile, "\tvals := []interface{}{}\n")
	appendToBytes(&findFile, "\tfor _, input := range inputs {\n")
	appendToBytes(&findFile, fmt.Sprintf("\t\tsqlStr += \"(%s),\"\n", preFieldsForInsert))
	appendToBytes(&findFile, fmt.Sprintf("\t\tvals = append(vals, %s)\n", modelFieldsForInsert))
	appendToBytes(&findFile, "\t}\n")
	appendToBytes(&findFile, "\tsqlStr = sqlStr[0 : len(sqlStr)-1]\n")
	appendToBytes(&findFile, "\tstmt, errStmt := db.GetDB().Prepare(sqlStr)\n")
	appendToBytes(&findFile, "\tdefer stmt.Close()\n")

	appendToBytes(&findFile, "\tif errStmt != nil {\n")
	appendToBytes(&findFile, "\t\tfmt.Println(\"errStmt = \", errStmt)\n")
	appendToBytes(&findFile, "\t\treturn 0, 0, errStmt\n")
	appendToBytes(&findFile, "\t}\n\n")
	appendToBytes(&findFile, "\tres, errInsert := stmt.Exec(vals...)\n")
	appendToBytes(&findFile, "\tif errInsert != nil {\n")
	appendToBytes(&findFile, "\t\tfmt.Println(\"errInsert = \", errInsert)\n")
	appendToBytes(&findFile, "\t\treturn 0, 0, errInsert\n")
	appendToBytes(&findFile, "\t}\n\n")
	appendToBytes(&findFile, "\trnoOfRows, _ := res.RowsAffected()\n")
	appendToBytes(&findFile, "\trlastId, _ := res.LastInsertId()\n")
	appendToBytes(&findFile, "\treturn rnoOfRows, rlastId, err\n")
	appendToBytes(&findFile, "}\n")
	///////////////////////////////
	errFindFile := ioutil.WriteFile(folderName+"/create.go", findFile, 0644)
	if errFindFile != nil {
		c.Println("Error while writing to file err = ", errFindFile)
	}

}
