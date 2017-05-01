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
	//root:root@tcp(192.168.40.11:3306)/sakila2
	db, err := sql.Open("mysql", "root:root@/sakila2?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	database = db
}

//GetDB ...
func GetDB() *sql.DB {
	return database
}
