package main

import (
	"runtime"
	"time"

	"github.com/ankitforcode/core-api/config"
	. "github.com/ankitforcode/core-api/lib/routes"
	"github.com/ankitforcode/core-api/lib/server"
	"github.com/gorilla/mux"
	log "github.com/inconshreveable/log15"
	mgo "gopkg.in/mgo.v2"
)

const layout = "2006/01/02 - 15:04:05"

func handleRequest() {
	//	router = routes.RegisterRoute()
	router := mux.NewRouter()
	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/single", ReturnArticle).Methods("GET")
	server.Run(router)
}

func init() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Info("Running with :", "CPU", nuCPU)
	config.Load()
	session, err := mgo.DialWithTimeout("localhost", 5*time.Second)
	if err != nil {
		log.Error("Error Connecting to Mongo : ", "err", err)
	} else {
		log.Info("Successfully Connected to Database Server :", "server", "localhost")
		anotherSession := session.Copy()
		defer anotherSession.Close()
	}
}

func main() {
	handleRequest()
}
