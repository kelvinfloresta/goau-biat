package commands

import (
	"goau-biat/config"
	"log"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func SmarRune(ev hook.Event) {
	if ev.Keychar == config.AttackRuneKey {
		delay := time.Millisecond * time.Duration(rand.Intn(100)+200)
		time.Sleep(delay)
		robotgo.Click()
		log.Default().Printf("Smart Rune Delay: %s", delay)
	}
}
