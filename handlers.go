package main

import (
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/AEPi-AK/game-server/models"
)


type ErrorResponse struct {
	Error string `json:"error"`
}

type HelloResponse struct {
	State models.State `json:"state"`
}

type AttackResponse struct {
 	State models.State `json:"state"`
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

	var requestData models.Poll 
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

	if err := r.Body.Close(); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	log.Info(requestData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

 	pollResponse := PollResponse{CanAttack: false, State: state}
	if err := json.NewEncoder(w).Encode(pollResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}
