package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"goau-biat/util/winapi"
	"math/rand"
	"sync"
	"time"

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

	ChangeWindow()
	util.Logger.Println("Making rune")
	for i := 0; i < 4; i++ {
		winapi.KeyTap(config.CreateRune)
		r.runeCount += 1
		util.Logger.Printf("%d Rune made\n", r.runeCount)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)

		if paused {
			return
		}
	}

	winapi.KeyTap("r")
	time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)

	r.restoreCurrentWindow()
}

func (r *RuneMaker) equipSoft() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	ChangeWindow()
	util.Logger.Println("Equipping Soft")
	for i := 0; i < 2; i++ {
		winapi.KeyTap(config.EquipSoft)
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

	ChangeWindow()
	util.Logger.Println("Equipping Ring")
	for i := 0; i < 2; i++ {
		winapi.KeyTap(config.EquipRing)
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

	ChangeWindow()
	util.Logger.Println("Staying logged in")
	winapi.KeyTap("up")
	util.RandomSleep(5, time.Second)
	winapi.KeyTap("down")
	util.RandomSleep(3, time.Second)
	winapi.KeyTap("right")
	util.RandomSleep(6, time.Second)
	winapi.KeyTap("left")
	r.restoreCurrentWindow()
}

func (r *RuneMaker) eatFood() {
	r.storeCurrentWindow()
	if paused {
		return
	}

	ChangeWindow()
	util.Logger.Println("Eating food")
	for i := 0; i < 4; i++ {
		winapi.KeyTap(config.EatFood)
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
	newCurrentWindow := winapi.GetPid()

	util.Logger.Printf("total process %d\n", r.processCount)
	if clientPid == newCurrentWindow || newCurrentWindow == r.currentWindow {
		util.Logger.Println("No pid was stored", r.processCount)
		return
	}

	util.Logger.Println("Storing pid")
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
	util.Logger.Printf("total process %d\n", r.processCount)

	if r.processCount > 0 {
		util.Logger.Println("window not restored. there is at least one process running")
		return
	}

	util.Logger.Println("Restoring the window")
	winapi.ActivePid(r.currentWindow)
}
