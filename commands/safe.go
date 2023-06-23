package commands

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func Safe(description string, ev hook.Event, key rune, safeKeys ...string) {
	if ev.Keychar != key {
		return
	}
	for _, safeKey := range safeKeys {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+150)
		time.Sleep(delay)
		robotgo.KeyTap(safeKey)
		logger.Printf("Safe %s: %s", description, delay)
	}
}
