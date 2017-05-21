package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Firstname string        `json:"firstname,omitempty"`
	Lastname  string        `json:"lastname,omitempty"`
	Address   *Address      `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	people = append(people, Person{Id: "1", Firstname: "Ankit", Lastname: "Agarwal", Address: &Address{City: "Dublin", State: "CA"}})
	json.NewEncoder(w).Encode(people)
	end := time.Now()
	latency := end.Sub(start)
	log.Info("Endpoint Hit: getpeople.", "Latency", latency, "Path", r.URL.Path, "Method", r.Method)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "Welcome to the HomePage!")
	end := time.Now()
	latency := end.Sub(start)
	log.Info("Endpoint Hit: homepage.", "Latency", latency, "Path", r.URL.Path, "Method", r.Method)
}

func ReturnArticle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "return a specific article")
	end := time.Now()
	latency := end.Sub(start)
	log.Info("Endpoint Hit: returnArticle.", "Latency", latency, "Path", r.URL.Path, "Method", r.Method)
}

func RedirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}
