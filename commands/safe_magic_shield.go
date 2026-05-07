package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"goau-biat/util/winapi"
	"math/rand"
	"time"

	hook "github.com/robotn/gohook"
)

func SafeMagicShield(ev hook.Event) {
	if ev.Keychar == config.SafeMagicShieldKey {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+50)
		time.Sleep(delay)
		winapi.KeyTap(config.ExuraVitaKey)
		util.Logger.Printf("Safe magic shield: %s", delay)
	}
}
