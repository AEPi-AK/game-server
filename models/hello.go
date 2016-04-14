package models

type Hello struct {
	Player GameEntity `json:"player"`
	PlayerNum int `json:"player_number"`
}
