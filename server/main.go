package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/myjwt"
	route "bitbucket.org/restapi/server/routes"
)

var port string = "8080"
var serviceName string = "onlinebooking/v1"

func PrimaryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello path:"+r.URL.Path)
	return
}

func handleCtrlC(c chan os.Signal) {
	sig := <-c
	// handle ctrl+c event here
	// for example, close database
	fmt.Println("\nsignal: ", sig)
	requestServer("removeServer")
	os.Exit(0)
}

func requestServer(action string) {
	///Register service///
	registerUrl := "http://localhost:9090/" + action + "?serviceName=" + serviceName + "&serverPath=localhost:" + port
	log.Println(" registerUrl = ", registerUrl)
	req, err := http.NewRequest("GET", registerUrl, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	////////End Register////////
}

func main() {
	////Initial Mysql, Redis, JWT
	logger.InitLogger()
	db.InitMysql()
	db.InitRedis()
	myjwt.InitKeys()
	defer db.GetDB().Close()

	////Ctrl+C handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c)

	////Register route for the server
	//route.NewRouter()
	http.Handle("/", route.NewRouter())

	////Config and start the server
	server := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	requestServer("addServer")
	server.ListenAndServe()
	//http.ListenAndServe(":8080", nil)

	// mux := bone.New()
	//
	// // mux.Get, Post, etc ... takes http.Handler
	//
	// // GetFunc, PostFunc etc ... takes http.HandlerFunc
	// mux.GetFunc("/test", middleware.Logger(PrimaryHandler, "test"))
	//
	// http.ListenAndServe(":8080", mux)

}
