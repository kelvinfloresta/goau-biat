package commands

import (
	"goau-biat/config"
	"os"

	hook "github.com/robotn/gohook"
)

func Exit(ev hook.Event) {
	if ev.Keychar == config.CancelRunes {
		os.Exit(0)
	}
}
