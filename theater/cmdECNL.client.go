package theater

import (
	"../GameSpy"
	"../log"
)

// ECNL - CLIENT calls when they want to leave
func (tM *TheaterManager) ECNL(event GameSpy.EventClientFESLCommand) {
	if !event.Client.IsActive {
		log.Noteln("Client left")
		return
	}

	//wantsToLeaveQueue = true

	answer := make(map[string]string)
	answer["TID"] = event.Command.Message["TID"]
	answer["GID"] = event.Command.Message["GID"]
	answer["LID"] = event.Command.Message["LID"]
	event.Client.WriteFESL("ECNL", answer, 0x0)
	tM.logAnswer("ECNL", answer, 0x0)
}
