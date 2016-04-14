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
{type: "hello", "player": player}
Adds a player, returns state

#### Attack
POST /attack
{type: "attack", "damage": int, "id": id}
Does an attack of "damage" to the identifier, returns the state.


#### Poll
POST /poll
{type: "poll", "id": identifier}
Returns {"can_attack": true/false, "state": state}, true if it is ready to take another attack, otherwise false.
