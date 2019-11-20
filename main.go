package main

import (
	"log"
	"net/http"

	"github.com/NextGenConf/SessionService/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	env := handlers.InitializeEnvironment()
	r.HandleFunc("/api/session/", env.GetAllSession).Methods("GET")
	r.HandleFunc("/api/session/{UniqueName}", env.GetSession).Methods("GET")
	r.HandleFunc("/api/session/", env.AddNewSession).Methods("POST")

	log.Print("Running and listening on port 5000")
	http.ListenAndServe(":5000", r)
}
