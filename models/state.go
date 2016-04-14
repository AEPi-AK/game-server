package models

type State struct {
	Player1  GameEntity `json:"player1"`
	Player2  GameEntity `json:"player2"`
	Monster GameEntity `json:"monster"`
}
