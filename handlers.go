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

func Hello(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("Received hello request")

	var requestData models.Hello 
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

 	helloResponse := PerformHello(requestData)
	if err := json.NewEncoder(w).Encode(helloResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

func HelloMonster(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("Received hello monster request")

	var requestData models.HelloMonster
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

 	helloMonsterResponse := PerformHelloMonster(requestData)
	if err := json.NewEncoder(w).Encode(helloMonsterResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

func Attack(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("Received attack request")

	var requestData models.Attack 
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

 	attackResponse := PerformAttack(requestData)
	if err := json.NewEncoder(w).Encode(attackResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

 	pollResponse := PerformPoll(requestData)
	if err := json.NewEncoder(w).Encode(pollResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}
