package models

type GameEntity struct {
	ID string `json:"id"`
	Hitpoints int `json:"hitpoints"`
	LastAttackUsed string `json:"last_attack"`
}
