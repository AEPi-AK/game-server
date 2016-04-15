#!/bin/bash

# Well. here we go. a journey through curl. 
echo add monster
curl -H "Content-Type: application/json" -d '{"monster": {"id": "scary monster", "hitpoints": 100}}' http://localhost:8000/hello-monster
echo add player 2
curl -H "Content-Type: application/json" -d '{"player": {"id": "jordan", "hitpoints": 100}, "player_number": 2}' http://localhost:8000/hello
echo add player 1
curl -H "Content-Type: application/json" -d '{"player": {"id": "everi", "hitpoints": 100}, "player_number": 1}' http://localhost:8000/hello

# Ensure BEN is not added

echo Both players should be able to attack, the monster shouldnt be able to attack
curl -H "Content-Type: application/json" -d '{"id": "jordan"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "everi"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "scary monster"}' http://localhost:8000/poll

echo Player 1 attack
curl -H "Content-Type: application/json" -d '{"attacker":"jordan", "target": "scary monster", "damage": 10, "attack_name": "slash"}' http://localhost:8000/attack
echo Player 1 should not be able to attack
curl -H "Content-Type: application/json" -d '{"id": "jordan"}' http://localhost:8000/poll
echo But player 2 should, monster still shouldnt
curl -H "Content-Type: application/json" -d '{"id": "everi"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "scary monster"}' http://localhost:8000/poll
echo Have player 2 attack. Now monster should be able to attack and both players shouldnt
curl -H "Content-Type: application/json" -d '{"attacker":"everi", "target": "scary monster", "damage": 10, "attack_name": "stab"}' http://localhost:8000/attack
curl -H "Content-Type: application/json" -d '{"id": "jordan"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "everi"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "scary monster"}' http://localhost:8000/poll

echo Have the monster attack
curl -H "Content-Type: application/json" -d '{"attacker":"scary monster", "target": "everi", "damage": 50}' http://localhost:8000/attack
echo Now monster shouldnt be able to attack but players should
curl -H "Content-Type: application/json" -d '{"id": "jordan"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "everi"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "scary monster"}' http://localhost:8000/poll
