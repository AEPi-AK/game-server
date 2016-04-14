package main

import (
	"strings"
	"sync"

	"github.com/AEPi-AK/game-server/models"
)

func PerformAttack(attack models.Attack) models.State {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	
	if (strings.EqualFold(state.Player1.ID, attack.Target)) {
		state.Player1.Hitpoints = state.Player1.Hitpoints - attack.Damage
	} else if (strings.EqualFold(state.Player2.ID, attack.Target)) {
		state.Player2.Hitpoints = state.Player2.Hitpoints - attack.Damage
	} else if (strings.EqualFold(state.Player2.ID, attack.Target)) {
		state.Monster.Hitpoints = state.Monster.Hitpoints - attack.Damage
	}

	mutex.Unlock()
	return state
}

func PerformHello(hello models.Hello) models.State {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	
	if (strings.EqualFold(state.Player1.ID, "")) {
		state.Player1 = hello.Player
	} else if (strings.EqualFold(state.Player2.ID, "")) {
		state.Player2 = hello.Player
	}

	mutex.Unlock()
	return state
}

func PerformPoll(poll models.Poll) PollResponse {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	isMonster := false

	if strings.EqualFold(poll.ID, state.Monster.ID) {
		isMonster = true
	} 

	canAttack := false
	if isMonster && monsterTurn {
		canAttack = true
		isMonster = false
		player1Attacked = false
		player2Attacked = false
	} else if !isMonster && (strings.EqualFold(poll.ID, state.Player1.ID) && !player1Attacked) {
		canAttack = true
		player1Attacked = true
	} else if !isMonster && (strings.EqualFold(poll.ID, state.Player2.ID) && !player2Attacked) {
		canAttack = true
		player2Attacked = true
	}

	mutex.Unlock()

	return PollResponse{CanAttack: canAttack, State: state}
}

func PerformHelloMonster(hello-monster models.HelloMonster) models.State {
	return state
}
