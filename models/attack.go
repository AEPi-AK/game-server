package models

type Attack struct {
	Attacker string `json:"attacker"`
	Target string `json:"target"`
	Damage int `json:"damage"`
}
