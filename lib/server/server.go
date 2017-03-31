package server

import (
	"fmt"
	log "github.com/inconshreveable/log15"
	"net/http"
)

type Server struct {
	Host	 	string 	`json:"Host"`  // Server name
	Port  		uint    `json:"Port"`  // HTTP port
}

func Run(s Server) {
	log.Info("Srarting Server On:" ,"address", s.Host, "port" , s.Port)
	startHTTP(s)
}
	
func startHTTP(s Server) {
	if err := http.ListenAndServe(httpAddress(s), nil); err != nil {
		log.Error("Error", "error", err)
	}
}

func httpAddress(s Server) string {
	return s.Host + ":" + fmt.Sprintf("%d", s.Port)
}
