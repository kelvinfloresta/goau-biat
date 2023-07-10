package commands

import (
	"goau-biat/config"
	"goau-biat/sounds"
	"goau-biat/util"
	"log"
	"sync"

	hook "github.com/robotn/gohook"
)

var paused = false
var pauseMu sync.Mutex

func togglePause(ev hook.Event) bool {
	if !isPauseEvent(ev) {
		return false
	}

	pauseMu.Lock()
	paused = !paused
	pauseMu.Unlock()

	log.Default().Printf("Paused: %t", paused)

	if paused {
		sounds.PlaySound(sounds.Paused)
	} else {
		sounds.PlaySound(sounds.Resumed)
	}

	return true
}

func isPauseEvent(ev hook.Event) bool {
	if !util.IsKeyDown(ev.Kind) {
		return false
	}

	return ev.Keychar == config.Pause || rune(ev.Rawcode) == config.Pause
}
