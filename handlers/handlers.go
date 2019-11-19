package handlers

import (
    "fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func GetSession(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I like %s!", mux.Vars(r)["key"])
}