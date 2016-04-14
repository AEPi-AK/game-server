package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/rs/cors"
	"github.com/AEPi-AK/game-server/models"
)

var (
	state models.State
	monsterTurn bool
	player1Attacked bool
	player2Attacked bool
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
