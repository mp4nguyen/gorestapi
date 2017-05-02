package userMdl

import (
	"database/sql"
	"fmt"

	"bitbucket.org/restapi/db"
)

func Create(user User) (userReturn sql.Result, err error) {

	sql := "INSERT INTO test_users set user_nickname='" + user.Name + "', user_first='" + user.First + "', user_last='" + user.Last + "', user_email='" + user.Email + "'"
	q, err := db.GetDB().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	return q, err
}
