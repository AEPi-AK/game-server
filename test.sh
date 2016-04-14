#!/bin/bash

# Well. here we go. a journey through curl. 
echo add monster
curl -H "Content-Type: application/json" -d '{"monster": {"id": "scary monster", "hitpoints": 100, "defense": 10}}' http://localhost:8000/hello-monster
echo add player 1
curl -H "Content-Type: application/json" -d '{"player": {"id": "jordan", "hitpoints": 100, "defense": 10}}' http://localhost:8000/hello
echo add player 2
curl -H "Content-Type: application/json" -d '{"player": {"id": "everi", "hitpoints": 100, "defense": 10}}' http://localhost:8000/hello

# Ensure BEN is not added
echo If ben gets added something went horribly wrong
curl -H "Content-Type: application/json" -d '{"player": {"id": "ben", "hitpoints": 100, "defense": 10}}' http://localhost:8000/hello

echo Both players should be able to attack
curl -H "Content-Type: application/json" -d '{"id": "jordan"}' http://localhost:8000/poll
curl -H "Content-Type: application/json" -d '{"id": "everi"}' http://localhost:8000/poll
echo The monster shouldnt be able to attack
curl -H "Content-Type: application/json" -d '{"id": "scary monster"}' http://localhost:8000/poll
