package utils

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/restapi/logger"
	"golang.org/x/crypto/bcrypt"
)

func LogError(info string, err error) {
	if err != nil {
		logger.Log.Errorf("%s (err = %s)", info, err)
	}
}

func ErrorHandler(info string, err error, tx *sql.Tx) {
	if err != nil {
		logger.Log.Errorf("%s (err = %s)", info, err)
		if tx != nil {
			tx.Rollback()
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	ErrorHandler("Check password", err, nil)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetValueOfField(f reflect.Value) string {
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(f.Int(), 10)
	case reflect.String:
		return f.String()
	case reflect.Struct:
		if f.Type() == reflect.TypeOf(time.Time{}) {
			return f.Interface().(time.Time).UTC().Format("2006-01-02 15:04:05")
		} else {
			return "Cannot get value of " + f.Type().String()
		}
	default:
		return "Cannot get value of " + f.Type().String()
	}
}

func BuildUpdateSQLString(tableName string, inf interface{}) (sqlStr string, sqlVals []interface{}) {
	t := reflect.TypeOf(inf)
	v := reflect.ValueOf(inf)
	tagName := "mysql"
	// Iterate over all available fields and read the tag value

	sqlString := " UPDATE " + tableName + " SET "
	whereString := " WHERE "
	whereValue := ""
	vals := []interface{}{}

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		fieldValue := v.Field(i)
		// Get the field tag value
		tag := field.Tag.Get(tagName)
		indexOfPRI := strings.Index(tag, "PRI")
		fmt.Printf("%d. %v (%v), tag: '%v'  value = %s\n", i+1, field.Name, field.Type.Name(), tag, GetValueOfField(fieldValue))
		if indexOfPRI > 0 {
			whereString += tag[0:indexOfPRI-1] + " = ? "
			whereValue = GetValueOfField(fieldValue)
		} else {
			sqlString += tag + " = ?,"
			vals = append(vals, GetValueOfField(fieldValue))
		}

	}

	sqlString = sqlString[0:len(sqlString)-1] + whereString
	vals = append(vals, whereValue)

	fmt.Println(" sql string = ", sqlString)
	fmt.Println(" sql vals = ", vals)
	return sqlString, vals
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
