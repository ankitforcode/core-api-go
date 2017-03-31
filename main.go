package main

import (
	"fmt"
	"encoding/json"
	"net/http"
        "time"
	"runtime"
	"os"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"./config"
//	"./config/routes"
        log "github.com/inconshreveable/log15"
)
	
const layout = "2006/01/02 - 15:04:05"

type Person struct {	
	ID		string 		`json:"id,omitempty"`
	Firstname	string		`json:"firstname,omitempty"`
	Lastname	string		`json:"lastname,omitempty"`
	Address		*Address	`json:"address,omitempty"`
}

type Address struct {
	City 		string		`json:"city"`
	State		string		`json:"state"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	json.NewEncoder(w).Encode(people)
        end := time.Now()
        latency := end.Sub(start)
        log.Info("Endpoint Hit: getpeople.", "Latency" , latency , "Path" , r.URL.Path , "Method" , r.Method)
}	

func homePage(w http.ResponseWriter, r *http.Request){
	start := time.Now()
	fmt.Fprintf(w, "Welcome to the HomePage!")
	end := time.Now()
	latency := end.Sub(start)
 	log.Info("Endpoint Hit: homepage.", "Latency" , latency , "Path" , r.URL.Path , "Method" , r.Method )
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "return a specific article")
	end := time.Now()
	latency := end.Sub(start)
	log.Info("Endpoint Hit: returnArticle.", "Latency" , latency , "Path" , r.URL.Path , "Method" , r.Method )
}

func handleRequest() {
	router := mux.NewRouter()
    	people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
   	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
	var listenAddress = fmt.Sprintf("0.0.0.0:%d", config.Config.Port)
	log.Info("Starting Core API Server On :", "address", listenAddress)
	router.HandleFunc("/", homePage).Methods("GET")
    	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/single", returnArticle).Methods("GET")
	err := http.ListenAndServe(listenAddress, router)
	// err := http.ListenAndServe(listenAddress, http.HandlerFunc(routes.RedirectToHTTPS))
	if err != nil {
		log.Error("ListenAndServe: ", "err", err)
		os.Exit(1)
	}
}

func main() {
	nuCPU := runtime.NumCPU()
	log.Info("Running with :", "CPU", nuCPU )
	var connection = fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Name)
	db, err := sql.Open("mysql", connection)
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Error("Error : ", "err", err.Error())
	} else {
		log.Info("Successfully Connected to Database Server: ", "server", config.Config.DB.Host, "database", config.Config.DB.Name)
	}
	runtime.GOMAXPROCS(nuCPU)
	handleRequest()
}
