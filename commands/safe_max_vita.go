package commands

import (
	"goau-biat/config"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func SafeMaxVita(ev hook.Event) {
	if ev.Keychar == config.SafeExuraMaxVitaKey {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+100)
		time.Sleep(delay)
		robotgo.KeyTap(config.SecondaryExuraVitaKey)
		logger.Printf("Safe exura max vita: %s", delay)
	}
}
