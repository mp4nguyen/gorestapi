package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/models/tableColumnsMdl"
	"github.com/abiosoft/ishell"
	logrus "github.com/sirupsen/logrus"
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

var log = logrus.New()

func main() {
	// Example of redirecting log output to a new file at runtime

	log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	//log.Formatter = &logrus.JSONFormatter{}
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	/////////////////
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

			// output, _ := json.Marshal(tableColumns.TableColumns)
			// fmt.Println(string(output))

			queryFields := ""
			outputFields := ""
			modelFields := ""
			preFieldsForInsert := ""
			modelFieldsForInsert := ""
			columnKey := ""
			declareDateFields := ""
			assignDateFields := ""
			for index, column := range tableColumns.TableColumns {
				queryFields = queryFields + "," + column.COLUMNNAME

				modelFields = modelFields + "" + fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName(column.COLUMNNAME), dataType(column.DATATYPE), jsonName(column.COLUMNNAME))
				if dataType(column.DATATYPE) == "time.Time" {
					modelFieldsForInsert = modelFieldsForInsert + ", input." + fieldName(column.COLUMNNAME) + ".Format(\"2006-01-02 15:04:05\")"

					declareDateFields += "temp" + fieldName(column.COLUMNNAME) + " := mysql.NullTime{} \n"
					outputFields = outputFields + ",&temp" + fieldName(column.COLUMNNAME)
					assignDateFields += "row." + fieldName(column.COLUMNNAME) + " = " + "temp" + fieldName(column.COLUMNNAME) + ".Time \n"
				} else {
					modelFieldsForInsert = modelFieldsForInsert + ", input." + fieldName(column.COLUMNNAME)
					outputFields = outputFields + ",&row." + fieldName(column.COLUMNNAME)
				}

				preFieldsForInsert = preFieldsForInsert + ",?"

				if column.COLUMNKEY == "PRI" {
					columnKey = column.COLUMNNAME
				}

				c.Println(index, ": ", column.COLUMNNAME, " ", column.DATATYPE, " ", column.CHARACTERMAXIMUMLENGTH, " ", column.COLUMNKEY)
			}
			queryFields = queryFields[1:len(queryFields)]
			outputFields = outputFields[1:len(outputFields)]
			modelFieldsForInsert = modelFieldsForInsert[1:len(modelFieldsForInsert)]
			preFieldsForInsert = preFieldsForInsert[1:len(preFieldsForInsert)]
			///// create folder for model////
			folderName := "outputs/" + LcFirst(modelName) + "Mdl"
			os.Mkdir(folderName, 0777)
			folderCtrlName := "outputs/" + LcFirst(modelName) + "Ctrl"
			os.Mkdir(folderCtrlName, 0777)

			createModelFile(c, folderName, modelName, modelFields)
			createFindFile(c, folderName, schemaName, tableName, modelName, outputFields, queryFields, declareDateFields, assignDateFields)
			createMapFindFile(c, folderName, schemaName, tableName, modelName, outputFields, queryFields, declareDateFields, assignDateFields, fieldName(columnKey))
			createFindByIdFile(c, folderName, schemaName, tableName, modelName, outputFields, queryFields, columnKey, declareDateFields, assignDateFields)
			createCreateFile(c, folderName, schemaName, tableName, modelName, preFieldsForInsert, modelFieldsForInsert, queryFields)
			//createFindCtrlFile(c, folderCtrlName, modelName)

		},
	})

	// simulate an authentication
	shell.AddCmd(&ishell.Cmd{
		Name: "relationship",
		Help: "Make relationship between models",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			detailTableForeignKey := ""
			masterTableForeignKey := ""
			detailColumnKey := ""
			masterColumnKey := ""

			c.Print("Make relationship between models")
			c.Print("Relationship type[1:one-one|2:one-many]: ")
			relationshipType := c.ReadLine()
			c.Print("Schema name: ")
			schemaName := c.ReadLine()
			c.Print("Master table name: ")
			masterTableName := c.ReadLine()
			c.Print("Relationship name: ")
			relationShipName := c.ReadLine()
			if relationshipType == "1" {
				c.Print("Master table foreign key: ")
				masterTableForeignKey = c.ReadLine()
			}
			c.Print("Master model name: ")
			masterModelName := UcFirst(c.ReadLine())
			c.Print("Detail table name: ")
			detailTableName := c.ReadLine()
			if relationshipType == "2" {
				c.Print("Master table foreign key: ")
				detailTableForeignKey = c.ReadLine()
			}
			c.Print("Detail model name: ")
			detailModelName := UcFirst(c.ReadLine())

			folderName := LcFirst(masterModelName) + "Mdl"

			masterTableColumns, err := tableColumnsMdl.GetTableColumn(schemaName, masterTableName)
			if err != nil {
				c.Printf("Cannot get columns for table %s.%s (%s) \n", schemaName, masterTableName, err)
				return
			}

			if len(masterTableColumns.TableColumns) == 0 {
				c.Printf("There are no columns in the table %s.%s \n", schemaName, masterTableName)
				return
			}

			detailTableColumns, err2 := tableColumnsMdl.GetTableColumn(schemaName, detailTableName)
			if err2 != nil {
				c.Printf("Cannot get columns for table %s.%s (%s) \n", schemaName, detailTableName, err2)
				return
			}

			if len(detailTableColumns.TableColumns) == 0 {
				c.Printf("There are no columns in the table %s.%s \n", schemaName, detailTableName)
				return
			}

			for _, column := range masterTableColumns.TableColumns {
				if column.COLUMNKEY == "PRI" {
					masterColumnKey = column.COLUMNNAME
				}
				//c.Println(index, ": ", column.COLUMNNAME, " ", column.DATATYPE, " ", column.CHARACTERMAXIMUMLENGTH, " ", column.COLUMNKEY)
			}

			for _, column := range detailTableColumns.TableColumns {
				if column.COLUMNKEY == "PRI" {
					detailColumnKey = column.COLUMNNAME
				}
				//c.Println(index, ": ", column.COLUMNNAME, " ", column.DATATYPE, " ", column.CHARACTERMAXIMUMLENGTH, " ", column.COLUMNKEY)
			}

			createModelRelationshipFile(c, folderName, masterModelName, relationshipType, relationShipName, detailModelName)
			createRelationshipFindFile(c, folderName, masterModelName, detailModelName, masterColumnKey, detailColumnKey, masterTableForeignKey, detailTableForeignKey, relationshipType, relationShipName)
			//fmt.Println("masterTableForeignKey = ", masterTableForeignKey)
		},
	})

	// start shell
	shell.Start()
}
