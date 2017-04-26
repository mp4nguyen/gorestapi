package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"bitbucket.org/restapi/models"
	"bitbucket.org/restapi/myjwt"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type UserController struct{}

type Token struct {
	Token string `json:"token"`
}

var usersModel = new(models.UsersModel)

func (u UserController) Login(w http.ResponseWriter, r *http.Request) {

	NewUser := models.User{}

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

	fmt.Fprintln(w, string(json))
}

func (u UserController) AfterLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "string(json)")
}

func (u UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	NewUser := models.User{}

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

	q, err := usersModel.Create(NewUser)
	fmt.Println(q)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Pragma", "no-cache")

	urlParams := mux.Vars(r)
	id := urlParams["id"]
	userId, err := strconv.ParseInt(id, 10, 64)

	data, err := usersModel.One(userId)
	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "No such user")
	case err != nil:
		log.Fatal(err)
	default:
		output, _ := json.Marshal(data)
		fmt.Fprintf(w, string(output))
	}
}

func (u UserController) UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("starting retrieval")
	start := 0
	limit := 10

	next := start + limit

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", "<http://localhost:8080/api/users?start="+string(next)+"; rel=\"next\"")

	data, err := usersModel.All()

	if err != nil {
		fmt.Println(err)
	}

	output, _ := json.Marshal(data)
	fmt.Fprintln(w, string(output))
}
