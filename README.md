# game-server
Game server for 2016 booth


## Communicating with the server
### State
The server will return the state of the game on most endpoints. The state will be a JSON object containing the data for all the players and the monster in the game. Each player and monster will be represented as a JSON object as well.

#### Player/Monster
{"id": identifier, "hitpoints": int, "defense": int}

#### State
{"player1" player1, "player2": player2, "monster": monster}

### Messages to the server
#### Hello
POST /hello
{"player": player, "player_number": 1/2}
Adds a player, returns state

#### Hello monster
POST /hello-monster
{"monster": monster}
Adds a monster, returns state. this state will be fresh (no characters other than the monsters).

#### Attack
POST /attack
{"damage": int, "attacker": attacker_id, "target": target_id}
Does an attack of "damage" to the identifier, returns the state.


#### Poll
POST /poll
{"id": identifier}
Returns {"can_attack": true/false, "state": state}, true if identifier can take another attack, otherwise false.
