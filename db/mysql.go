package db

import (
	"database/sql"
	"fmt"
	"log"

	"bitbucket.org/restapi/utils"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

//DB ...
type DB struct {
	*sql.DB
}

//Init ...
func InitMysql() {
	//root:root@tcp(192.168.40.11:3306)/sakila2
	//"pnguyen:root@tcp(192.168.40.11:3306)/sakila2?parseTime=true"
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ocs?parseTime=true")
	if err != nil {
		log.Fatal("mysql.go: failed to connect to mysql ", err)
	}

	database = db
}

//GetDB ...
func GetDB() *sql.DB {
	return database
}

func Transaction(txFn func(*sql.Tx) error) (err error) {
	tx, err := GetDB().Begin()
	if err != nil {
		utils.ErrorHandler("Failed to create mysql transaction", err, nil)
		return
	}
	defer func() {
		if err != nil {
			utils.ErrorHandler("rollback mysql transaction", err, nil)
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	err = txFn(tx)
	return err
}

func Update(tableName string, m interface{}, tx *sql.Tx) (noOfEffects int64, err error) {

	sqlString, sqlVals := utils.BuildUpdateSQLString(tableName, m)
	stmt, errStmt := GetDB().Prepare(sqlString)
	if tx != nil {
		stmt, errStmt = tx.Prepare(sqlString)
	}
	defer stmt.Close()

	if errStmt != nil {
		fmt.Println("errStmt = ", errStmt)
		return 0, errStmt
	}

	res, errInsert := stmt.Exec(sqlVals...)
	if errInsert != nil {
		fmt.Println("errInsert = ", errInsert)
		return 0, errInsert
	}

	rnoOfRows, _ := res.RowsAffected()

	return rnoOfRows, err

}
