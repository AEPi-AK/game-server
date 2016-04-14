package main

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/AEPi-AK/game-server/models"
)


type ErrorResponse struct {
	Error string `json:"error"`
}

type PollResponse struct {
	CanAttack bool `json:"can_attack"`
	State models.State `json:"state"`
}

// Given a message, responds with a JSON object containing that message 
// as an error string/
func RespondBadRequest(w http.ResponseWriter, message string) {
	log.WithFields(log.Fields{
		"time":    time.Now(),
		"message": message,
	}).Error("Received a bad request")

	errorResponse := ErrorResponse{Error: message}
	http.Error(w, "", http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(errorResponse)
}

func Poll(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("Received poll request")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

 	pollResponse := PollResponse{CanAttack: false, State: state}
	if err := json.NewEncoder(w).Encode(pollResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}
