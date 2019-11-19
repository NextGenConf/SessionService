package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetSession(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I like %s!", mux.Vars(r)["key"])
}
