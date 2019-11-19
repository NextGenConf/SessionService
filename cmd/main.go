package main

import (
	"net/http"
	"github.com/NextGenConf/SessionService/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/session/{key}", handlers.GetSession).Methods("GET")
	http.ListenAndServe(":5000", r)
}