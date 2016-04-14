package main

import (
	"encoding/json"
	"net/http"
	"time"
	"io"

	log "github.com/Sirupsen/logrus"
)


type ErrorResponse struct {
	Error string `json:"error"`
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
	// TODO complete me

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello world!")	

}
