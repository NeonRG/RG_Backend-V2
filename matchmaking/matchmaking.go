package matchmaking

import (
	//"crypto/tls"
	//"io/ioutil"
	//"net/http"
	//"strings"

	"../GameSpy"
	"../log"
)

// Games - a list of available games
var Games = make(map[string]*GameSpy.Client)

var Shard string

// FindAvailableGID - returns a GID suitable for the player to join (ADD A PID HERE)
func FindAvailableGIDs() string {
	log.Debugln("Call")
	var gameID string

	for k := range Games {
		gameID = k
		log.Debugln("Server:" + k)
	}

	return gameID
}