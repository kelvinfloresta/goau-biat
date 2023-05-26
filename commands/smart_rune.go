package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"log"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func SmartRune(ev hook.Event) {
	if ev.Rawcode == config.AttackRuneKey && util.IsKeyDown(ev.Kind) {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+200)
		time.Sleep(delay)
		robotgo.Click()
		log.Default().Printf("Smart Rune Delay: %s", delay)
	}
}
