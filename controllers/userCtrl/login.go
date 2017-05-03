package userCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/models/userMdl"
	"bitbucket.org/restapi/myjwt"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

func LoginJWT(w http.ResponseWriter, r *http.Request) {

	NewUser := userMdl.User{}

	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&NewUser); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	output, err := json.Marshal(NewUser)
	log.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	// Create the token

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = "test"
	token.Claims = claims

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		if err != nil {
			log.Fatal(err)
		}
	}

	tokenString, err := token.SignedString(myjwt.SignKey)

	fmt.Println(" token 1= ", tokenString, err)

	response := Token{tokenString}
	json, _ := json.Marshal(response)

	fmt.Println(json)

	fmt.Fprintln(w, string(json))
}

func LoginAT(w http.ResponseWriter, r *http.Request) {

	NewUser := userMdl.User{}

	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&NewUser); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	output, err := json.Marshal(NewUser)
	log.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	at, err := accessTokenMdl.Create(1)
	//atObject, err2 := accessTokenMdl.One(at)

	//fmt.Println("AccessToken = ", atObject, err2)
	fmt.Fprintln(w, string(at))
}
