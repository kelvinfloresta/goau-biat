package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"math/rand"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type RuneMaker struct {
	foodTicker          *time.Ticker
	antiLogoutTicker    *time.Ticker
	ringTicker          *time.Ticker
	softTicker          *time.Ticker
	makeRuneTicker      *time.Ticker
	runeCount           int
	currentWindow       int
	processCount        int
	processCountMutex   sync.Mutex
	ShouldRestoreWindow bool
}

func (r *RuneMaker) Start() {
	r.foodTicker = time.NewTicker(config.FoodTime)
	r.antiLogoutTicker = time.NewTicker(config.AntiLogoutTime)
	r.ringTicker = time.NewTicker(config.RingTime)
	r.softTicker = time.NewTicker(config.SoftTime)
	r.makeRuneTicker = time.NewTicker(config.RuneTime)

	go r.forceStart()

	go r.togglePauseRune()

	for {
		select {
		case <-r.foodTicker.C:
			util.RandomSleep(5, time.Second)
			r.eatFood()

		case <-r.antiLogoutTicker.C:
			util.RandomSleep(5, time.Second)
			r.stayLoggedIn()

		case <-r.ringTicker.C:
			util.RandomSleep(10, time.Second)
			r.equipRing()

		case <-r.softTicker.C:
			util.RandomSleep(10, time.Second)
			r.equipSoft()

		case <-r.makeRuneTicker.C:
			util.RandomSleep(10, time.Second)
			r.makeRune()
		}
	}
}

func (r *RuneMaker) makeRune() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	changeWindow()
	logger.Println("Making rune")
	for i := 0; i < 4; i++ {
		robotgo.KeyTap(config.CreateRune)
		r.runeCount += 1
		logger.Printf("%d Rune made\n", r.runeCount)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)

		if paused {
			return
		}
	}
	r.restoreCurrentWindow()
}

func (r *RuneMaker) equipSoft() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	changeWindow()
	logger.Println("Equipping Soft")
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipSoft)
		util.RandomSleep(2, time.Second)

		if paused {
			return
		}
	}
	r.restoreCurrentWindow()
}

func (r *RuneMaker) equipRing() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	changeWindow()
	logger.Println("Equipping Ring")
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipRing)
		util.RandomSleep(2, time.Second)

		if paused {
			return
		}
	}
	r.restoreCurrentWindow()
}

func (r *RuneMaker) stayLoggedIn() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	changeWindow()
	logger.Println("Staying logged in")
	robotgo.KeyTap("up")
	util.RandomSleep(5, time.Second)
	robotgo.KeyTap("down")
	util.RandomSleep(3, time.Second)
	robotgo.KeyTap("right")
	util.RandomSleep(6, time.Second)
	robotgo.KeyTap("left")
	r.restoreCurrentWindow()
}

func (r *RuneMaker) eatFood() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	changeWindow()
	logger.Println("Eating food")
	for i := 0; i < 4; i++ {
		robotgo.KeyTap(config.EatFood)
		util.RandomSleep(2, time.Second)

		if paused {
			return
		}
	}
	r.restoreCurrentWindow()
}

func (r *RuneMaker) forceStart() {
	r.equipRing()
	r.equipSoft()
	r.eatFood()
	r.makeRune()
}

func (r *RuneMaker) togglePauseRune() {
	events := hook.Start()
	for ev := range events {
		toggled := togglePause(ev)
		if toggled && paused {
			r.foodTicker.Stop()
			r.antiLogoutTicker.Stop()
			r.ringTicker.Stop()
			r.softTicker.Stop()
			r.makeRuneTicker.Stop()
			continue
		}

		if toggled && !paused {
			r.foodTicker.Reset(config.FoodTime)
			r.antiLogoutTicker.Reset(config.AntiLogoutTime)
			r.ringTicker.Reset(config.RingTime)
			r.softTicker.Reset(config.SoftTime)
			r.makeRuneTicker.Reset(config.RuneTime)
			go r.forceStart()
		}
	}
}

func (r *RuneMaker) storeCurrentWindow() {
	if !r.ShouldRestoreWindow {
		return
	}
	r.processCountMutex.Lock()
	r.processCount += 1
	r.processCountMutex.Unlock()
	newCurrentWindow := robotgo.GetPid()

	logger.Printf("total process %d\n", r.processCount)
	if clientPid == newCurrentWindow || newCurrentWindow == r.currentWindow {
		logger.Println("No pid was stored", r.processCount)
		return
	}

	logger.Println("Storing pid")
	r.currentWindow = newCurrentWindow
}

func (r *RuneMaker) restoreCurrentWindow() {
	if !r.ShouldRestoreWindow {
		return
	}

	time.Sleep(1 * time.Second)
	r.processCountMutex.Lock()
	r.processCount -= 1
	r.processCountMutex.Unlock()
	logger.Printf("total process %d\n", r.processCount)

	if r.processCount > 0 {
		logger.Println("window not restored. there is at least one process running")
		return
	}

	logger.Println("Restoring the window")
	robotgo.ActivePid(r.currentWindow)
}
