package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NextGenConf/SessionService/models"
	"github.com/gorilla/mux"
)

type Environment struct {
	db models.SessionDatabaseHandler
}

func InitializeEnvironment() *Environment {
	return &Environment{
		db: models.InitializeDatabaseHandler(),
	}
}

func (e *Environment) GetSession(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["UniqueName"]
	session, err := e.db.GetSession(param)
	if err != nil {
		log.Printf("Failed to get session from db: %s", err.Error())
		http.Error(w, "Failed to get session from database", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(session)
	if err != nil {
		log.Printf("Failed to serialize sessions: %s", err.Error())
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (e *Environment) GetAllSession(w http.ResponseWriter, r *http.Request) {
	sessions, err := e.db.GetAllSessions()
	if err != nil {
		http.Error(w, "Failed get response from db", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(sessions)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (e *Environment) AddNewSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		log.Printf("Failed to decode response: %s. Returning bad request.", err.Error())
		http.Error(w, "Wrongly formatted body", http.StatusBadRequest)
		return
	}

	if err := e.db.AddSession(session); err != nil {
		http.Error(w, "Failed to insert into db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
