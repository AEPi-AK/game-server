package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/rs/cors"
)

var (
	state models.State
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	handler := cors.Default().Handler(NewRouter())
	log.Info("Listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
