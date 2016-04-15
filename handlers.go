package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AEPi-AK/game-server/models"
	log "github.com/Sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type PollResponse struct {
	CanAttack bool         `json:"can_attack"`
	State     models.State `json:"state"`
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

	log.WithFields(log.Fields{
		"time": time.Now(),
		"data": helloResponse,
	}).Info("Received hello request")

	if err := json.NewEncoder(w).Encode(helloResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

func HelloMonster(w http.ResponseWriter, r *http.Request) {
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

	log.WithFields(log.Fields{
		"time": time.Now(),
		"data": helloMonsterResponse,
	}).Info("Received hello monster request")

	if err := json.NewEncoder(w).Encode(helloMonsterResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

func Attack(w http.ResponseWriter, r *http.Request) {
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

	log.WithFields(log.Fields{
		"time": time.Now(),
		"data": attackResponse,
	}).Info("Received attack request")

	if err := json.NewEncoder(w).Encode(attackResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

func Poll(w http.ResponseWriter, r *http.Request) {
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

	log.WithFields(log.Fields{
		"time":      time.Now(),
		"data":      pollResponse,
		"requester": requestData.ID,
	}).Info("Received poll request")

	if err := json.NewEncoder(w).Encode(pollResponse); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}
