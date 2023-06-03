package commands

import (
	"goau-biat/config"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func SafeMagicShield(ev hook.Event) {
	if ev.Keychar == config.SafeMagicShieldKey {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+100)
		time.Sleep(delay)
		robotgo.KeyTap(config.ExuraVitaKey)
		logger.Printf("Safe magic shield: %s", delay)
	}
}
