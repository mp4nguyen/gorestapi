package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

//DB ...
type DB struct {
	*sql.DB
}

//Init ...
func InitMysql() {

	db, err := sql.Open("mysql", "root:root@/OCS")
	if err != nil {
		log.Fatal(err)
	}

	database = db
}

//GetDB ...
func GetDB() *sql.DB {
	return database
}
