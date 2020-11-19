package fesl

import (
	//"encoding/binary"
	//"fmt"
	//"net"
	//"strconv"
	//"strings"

	"../GameSpy"
	"../log"
	"../matchmaking"
)

// Status - Basic fesl call to get overall service status (called before pnow?)
func (fM *FeslManager) Status(event GameSpy.EventClientTLSCommand) {
	if !event.Client.IsActive {
		log.Noteln("Client left")
		return
	}

	// Check if user is allowed to matchmake
	if !fM.userHasPermission(event.Client.RedisState.Get("uID"), "game.matchmake") {
		log.Noteln("User not worthy: " + event.Client.RedisState.Get("username"))
		fM.sendDenied(event)
		return
	}

	log.Noteln("STATUS CALLED")

	answer := make(map[string]string)
	answer["TXN"] = "Status"
	answer["id.id"] = "1"
	answer["id.partition"] = event.Command.Message["partition.partition"]
	answer["sessionState"] = "COMPLETE"
	answer["props.{}.[]"] = "2"
	answer["props.{resultType}"] = "JOIN"

	// Find latest game (do better later)
	//ipint := binary.BigEndian.Uint32(event.Client.IpAddr.(*net.TCPAddr).IP.To4())
	gameID := matchmaking.FindAvailableGIDs()

	answer["props.{games}.0.lid"] = "1"
	answer["props.{games}.0.fit"] = "1001"
	answer["props.{games}.0.gid"] = gameID
	answer["props.{games}.[]"] = "1"

	event.Client.WriteFESL("pnow", answer, 0x80000000)
	fM.logAnswer("pnow", answer, 0x80000000)
}

func (fM *FeslManager) sendDenied(event GameSpy.EventClientTLSCommand) {
	answer := make(map[string]string)
	answer["TXN"] = "Status"
	answer["id.id"] = "1"
	answer["id.partition"] = event.Command.Message["partition.partition"]
	answer["sessionState"] = "COMPLETE"
	answer["props.{}.[]"] = "2"
	answer["props.{resultType}"] = "JOIN"
	answer["props.{games}.[]"] = "0"
	event.Client.WriteFESL("pnow", answer, 0x80000000)
	fM.logAnswer("pnow", answer, 0x80000000)
}
