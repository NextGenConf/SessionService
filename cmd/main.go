package main

import (
	"github.com/NextGenConf/SessionService/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/session/{key}", handlers.GetSession).Methods("GET")
	http.ListenAndServe(":5000", r)
}
