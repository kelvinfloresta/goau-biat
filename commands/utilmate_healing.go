package commands

import (
	"goau-biat/config"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func UltimateHealing(ev hook.Event) {
	if ev.Keychar == config.Uh {
		randX := rand.Intn(20) + config.UhX
		randY := rand.Intn(3) + config.UhY

		delay := rand.Intn(100) + 100
		time.Sleep(time.Millisecond * time.Duration(delay))
		robotgo.MoveClick(randX, randY)
	}

	if ev.Keychar == config.UhRp {
		randX := rand.Intn(20) + config.UhX
		randY := rand.Intn(3) + config.UhY + 20

		delay := rand.Intn(100) + 100
		time.Sleep(time.Millisecond * time.Duration(delay))
		robotgo.MoveClick(randX, randY)
	}
}
