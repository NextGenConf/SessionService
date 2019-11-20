package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NextGenConf/SessionService/models"
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
	sessions := e.db.GetSessions()
	jsonData, err := json.Marshal(sessions)
	if err != nil {
		log.Printf("Failed to serialize sessions: %s", err.Error())
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (e *Environment) GetAllSession(w http.ResponseWriter, r *http.Request) {
	sessions := e.db.GetAllSessions()
	jsonData, err := json.Marshal(sessions)
	if err != nil {
		log.Printf("Failed to serialize sessions: %s", err.Error())
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
