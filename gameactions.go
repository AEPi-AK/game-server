package main

import (
	"github.com/AEPi-AK/game-server/models"
)

func PerformAttack(attack models.Attack) models.State {
	return state
}

func PerformHello(hello models.Hello) models.State {
	return state
}

func PerformPoll(poll models.Poll) PollResponse {
	return PollResponse{CanAttack: false, State: state}
}

func PerformHelloMonster(hello-monster models.HelloMonster) models.State {
	return state
}
