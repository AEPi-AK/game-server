package models

type GameEntity struct {
	ID string `json:"id"`
	Hitpoints int `json:"hitpoints"`
	Defense int `json:"defense"`
}