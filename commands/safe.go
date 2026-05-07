package commands

import (
	"goau-biat/util"
	"goau-biat/util/winapi"
	"math/rand"
	"time"

	hook "github.com/robotn/gohook"
)

func Safe(description string, ev hook.Event, key rune, safeKeys ...string) {
	if ev.Keychar != key {
		return
	}
	for _, safeKey := range safeKeys {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+150)
		time.Sleep(delay)
		winapi.KeyTap(safeKey)
		util.Logger.Printf("Safe %s: %s", description, delay)
	}
}
