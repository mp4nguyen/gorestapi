package userMdl

import "bitbucket.org/restapi/db"

func One(userId int64) (user User, err error) {

	ReadUser := User{}
	err = db.GetDB().QueryRow("select * from test_users where user_id=?", userId).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)

	return ReadUser, err
}
