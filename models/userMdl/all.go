package userMdl

import (
	"log"

	"bitbucket.org/restapi/db"
)

func All() (users Users, err error) {
	log.Println("starting retrieval all users")
	rows, err := db.GetDB().Query("select * from test_users LIMIT 10")
	if err != nil {
		log.Println("users.go: All() err = ", err)
	}

	Response := Users{}

	for rows.Next() {

		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)

		Response.Users = append(Response.Users, user)
	}

	return Response, err
}
