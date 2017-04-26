package models

import (
	"database/sql"
	"fmt"
	"log"

	"bitbucket.org/restapi/db"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Users struct {
	Users []User `json:"users"`
}

type UsersModel struct{}

func (m UsersModel) All() (users Users, err error) {
	log.Println("starting retrieval all users")
	rows, err := db.GetDB().Query("select * from test_users LIMIT 10")
	Response := Users{}

	for rows.Next() {

		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)

		Response.Users = append(Response.Users, user)
	}

	return Response, err
}

func (m UsersModel) One(userId int64) (user User, err error) {

	ReadUser := User{}
	err = db.GetDB().QueryRow("select * from test_users where user_id=?", userId).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)

	return ReadUser, err
}

func (m UsersModel) Create(user User) (userReturn sql.Result, err error) {

	sql := "INSERT INTO test_users set user_nickname='" + user.Name + "', user_first='" + user.First + "', user_last='" + user.Last + "', user_email='" + user.Email + "'"
	q, err := db.GetDB().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	return q, err
}
