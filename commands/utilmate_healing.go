package commands

import (
	"fmt"
	"goau-biat/config"
	"goau-biat/sounds"
	"goau-biat/util"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var pauseUh = false
var uhMu sync.Mutex

func togglePauseUh(ev hook.Event) bool {
	isPauseEvent := ev.Rawcode == config.PauseUh && util.IsKeyDown(ev.Kind)
	if !isPauseEvent {
		return false
	}

	fmt.Println(ev)
	uhMu.Lock()
	pauseUh = !pauseUh
	uhMu.Unlock()

	log.Default().Printf("Paused Uh: %t", pauseUh)

	if pauseUh {
		sounds.PlaySound(sounds.Paused)
	} else {
		sounds.PlaySound(sounds.Resumed)
	}

	return pauseUh
}

func UltimateHealing(ev hook.Event) {
	paused := togglePauseUh(ev)

	if paused {
		return
	}

	if ev.Rawcode == config.Uh && util.IsKeyDown(ev.Kind) {
		uhEk()
		robotgo.MilliSleep(20)
		robotgo.KeyTap("f6")
		uhEk()
	}

	if ev.Rawcode == config.UhRp && util.IsKeyDown(ev.Kind) {
		uhRp()
		robotgo.MilliSleep(20)
		robotgo.KeyTap("f6")
		uhRp()
	}
}

func uhRp() {
	randX := rand.Intn(20) + config.UhX
	randY := rand.Intn(3) + config.UhY + config.UhYDiff

	delay := time.Millisecond * time.Duration(rand.Intn(100)+100)
	time.Sleep(delay)
	forceClick(randX, randY)
	log.Default().Printf("UH RP Delay: %s", delay)
}

func uhEk() {
	randX := rand.Intn(20) + config.UhX
	randY := rand.Intn(3) + config.UhY

	delay := time.Millisecond * time.Duration(rand.Intn(100)+100)
	time.Sleep(delay)
	forceClick(randX, randY)
	log.Default().Printf("UH EK Delay: %s", delay)
}
