package commands

import (
	"goau-biat/config"
	"goau-biat/sounds"
	"log"

	hook "github.com/robotn/gohook"
)

var paused = false

func togglePause(ev hook.Event) bool {
	if !isPauseEvent(ev) {
		return false
	}

	paused = !paused
	log.Default().Printf("Paused: %t", paused)

	if paused {
		sounds.PlaySound(sounds.Paused)
	} else {
		sounds.PlaySound(sounds.Resumed)
	}

	return true
}

func isPauseEvent(ev hook.Event) bool {
	return ev.Keychar == config.Pause || rune(ev.Rawcode) == config.Pause
}
