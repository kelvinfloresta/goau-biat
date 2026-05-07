package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"goau-biat/util/winapi"
	"math/rand"
	"time"

	hook "github.com/robotn/gohook"
)

func SafeMaxVita(ev hook.Event) {
	if ev.Keychar == config.SafeExuraMaxVitaKey {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+50)
		time.Sleep(delay)
		winapi.KeyTap(config.SecondaryExuraVitaKey)
		util.Logger.Printf("Safe exura max vita: %s", delay)
	}
}
