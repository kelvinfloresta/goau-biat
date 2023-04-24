package commands

import (
	"goau-biat/config"
	"log"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func UltimateHealing(ev hook.Event) {
	if ev.Keychar == config.Uh {
		randX := rand.Intn(20) + config.UhX
		randY := rand.Intn(3) + config.UhY

		delay := time.Millisecond * time.Duration(rand.Intn(100)+100)
		time.Sleep(delay)
		robotgo.MoveClick(randX, randY)
		log.Default().Printf("UH EK Delay: %s", delay)
	}

	if ev.Keychar == config.UhRp {
		randX := rand.Intn(20) + config.UhX
		randY := rand.Intn(3) + config.UhY + 20

		delay := time.Millisecond * time.Duration(rand.Intn(100)+100)
		time.Sleep(delay)
		robotgo.MoveClick(randX, randY)
		log.Default().Printf("UH RP Delay: %s", delay)
	}
}
