package commands

import (
	"goau-biat/config"
	"log"

	hook "github.com/robotn/gohook"
)

var paused = false

func togglePause(ev hook.Event) bool {
	if isPauseEvent(ev) {
		paused = !paused
		log.Default().Printf("Paused: %t", paused)
		return true
	}

	return false
}

func isPauseEvent(ev hook.Event) bool {
	return ev.Keychar == config.Pause || rune(ev.Rawcode) == config.Pause
}
