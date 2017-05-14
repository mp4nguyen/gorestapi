package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/models/tableColumnsMdl"
	"github.com/abiosoft/ishell"
)

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func fieldName(mysqlFieldName string) string {
	names := strings.Split(mysqlFieldName, "_")
	returnName := ""
	for _, name := range names {
		returnName = returnName + UcFirst(name)
	}
	return returnName
}

func jsonName(mysqlFieldName string) string {
	names := strings.Split(mysqlFieldName, "_")
	returnName := ""
	for index, name := range names {
		if index == 0 {
			returnName = returnName + LcFirst(name)
		} else {
			returnName = returnName + UcFirst(name)
		}

	}
	return returnName
}

func dataType(mysqlDataType string) string {
	if mysqlDataType == "varchar" {
		return "string"
	} else if mysqlDataType == "int" || mysqlDataType == "tinyint" {
		return "int"
	} else if mysqlDataType == "datetime" {
		return "time.Time"
	} else {
		return mysqlDataType
	}

}

func appendToBytes(original *([]byte), addMore string) {
	*original = append(*original, addMore...)
}

func main() {

	db.InitMysql()
	defer db.GetDB().Close()
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Sample Interactive Shell")

	// register a function for "create model from table" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "model",
		Help: "create model for Golang",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			c.Print("The app will create model folder for Mysql table")
			c.Print("Schema name: ")
			schemaName := c.ReadLine()
			c.Print("Table name: ")
			tableName := c.ReadLine()
			c.Print("Model name: ")
			modelName := UcFirst(c.ReadLine())

			tableColumns, err := tableColumnsMdl.GetTableColumn(schemaName, tableName)
			if err != nil {
				c.Printf("Cannot get columns for table %s.%s (%s) \n", schemaName, tableName, err)
				return
			}

			if len(tableColumns.TableColumns) == 0 {
				c.Printf("There are no columns in the table %s.%s \n", schemaName, tableName)
				return
			}

			output, _ := json.Marshal(tableColumns.TableColumns)
			fmt.Println(string(output))

			queryFields := ""
			outputFields := ""
			///// create folder for model////
			folderName := LcFirst(modelName) + "Mdl"
			os.Mkdir(folderName, 0777)
			folderCtrlName := LcFirst(modelName) + "Ctrl"
			os.Mkdir(folderCtrlName, 0777)

			////create 0model.go////
			modelFile := []byte("")
			appendToBytes(&modelFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))
			appendToBytes(&modelFile, fmt.Sprintf("import \"time\"\n\n"))
			appendToBytes(&modelFile, fmt.Sprintf("type %s struct{\n", modelName))
			for index, column := range tableColumns.TableColumns {
				queryFields = queryFields + "," + column.COLUMNNAME
				outputFields = outputFields + ",&row." + fieldName(column.COLUMNNAME)
				c.Println(index, ": ", column.COLUMNNAME, " ", column.DATATYPE, " ", column.CHARACTERMAXIMUMLENGTH)
				appendToBytes(&modelFile, fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName(column.COLUMNNAME), dataType(column.DATATYPE), jsonName(column.COLUMNNAME)))
			}
			queryFields = queryFields[1:len(queryFields)]
			outputFields = outputFields[1:len(outputFields)]
			appendToBytes(&modelFile, fmt.Sprintf("\t}\n\n"))
			appendToBytes(&modelFile, fmt.Sprintf("type %ss []%s", modelName, modelName))
			errModelFile := ioutil.WriteFile(folderName+"/0model.go", modelFile, 0644)
			if errModelFile != nil {
				c.Println("Error while writing to file err = ", errModelFile)
			}
			////create find.go////
			findFile := []byte("")
			appendToBytes(&findFile, fmt.Sprintf("package %sMdl\n\n", LcFirst(modelName)))
			appendToBytes(&findFile, fmt.Sprintf("import \"time\"\n"))
			appendToBytes(&findFile, fmt.Sprintf("import \"log\"\n"))
			appendToBytes(&findFile, fmt.Sprintf("import \"bitbucket.org/restapi/db\"\n\n"))

			appendToBytes(&findFile, fmt.Sprintf("func Find()(%ss %ss,err error){\n", LcFirst(modelName), modelName))
			appendToBytes(&findFile, fmt.Sprintf("\trows, err := db.GetDB().Query(\"select %s from %s.%s\")\n", queryFields, schemaName, tableName))
			appendToBytes(&findFile, "\tif err != nil {\n")
			appendToBytes(&findFile, fmt.Sprintf("\t\tlog.Println(\"%sMdl.find.go: All() err = \", err)\n", LcFirst(modelName)))
			appendToBytes(&findFile, "\t}\n\n")
			appendToBytes(&findFile, fmt.Sprintf("\tresponse := %ss{}\n", modelName))
			appendToBytes(&findFile, "\tfor rows.Next() {\n")
			appendToBytes(&findFile, fmt.Sprintf("\t\trow := %s{}\n", modelName))
			appendToBytes(&findFile, fmt.Sprintf("\t\trows.Scan(%s)\n", outputFields))
			appendToBytes(&findFile, fmt.Sprintf("\t\tresponse = append(response,row)\n"))
			appendToBytes(&findFile, "\t}\n\n")
			appendToBytes(&findFile, "\treturn response, err\n")
			appendToBytes(&findFile, "}\n")

			errFindFile := ioutil.WriteFile(folderName+"/find.go", findFile, 0644)
			if errFindFile != nil {
				c.Println("Error while writing to file err = ", errFindFile)
			}

			////create findCtrl.go
			// calParams := GetCalendarParams{}
			// dec := json.NewDecoder(r.Body)
			// log.Println("dec Body = ", dec)
			// //fmt.Println(dec)
			// //fmt.Println(r.FormValue("id"))
			// for {
			// 	if err := dec.Decode(&calParams); err == io.EOF {
			// 		break
			// 	} else if err != nil {
			// 		log.Fatal(err)
			// 	}
			// }
			// output, err := json.Marshal(calParams)
			// log.Println(string(output))
			// if err != nil {
			// 	fmt.Println("Something went wrong!")
			// }
			////
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

			//			c.Printf("Authentication Successful. with username = %s and password = %s", username, password)
		},
	})

	// simulate an authentication
	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "simulate a login",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			// get username
			c.Print("Username: ")
			username := c.ReadLine()

			// get password.
			c.Print("Password: ")
			password := c.ReadPassword()

			c.Printf("Authentication Successful. with username = %s and password = %s", username, password)
		},
	})

	// start shell
	shell.Start()
}
