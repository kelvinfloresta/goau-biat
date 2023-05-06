package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	foodTicker       *time.Ticker
	antiLogoutTicker *time.Ticker
	ringTicker       *time.Ticker
	softTicker       *time.Ticker
	makeRuneTicker   *time.Ticker
	runeCount        = 0
)

func StartRune() {
	foodTicker = time.NewTicker(config.FoodTime)
	antiLogoutTicker = time.NewTicker(config.AntiLogoutTime)
	ringTicker = time.NewTicker(config.RingTime)
	softTicker = time.NewTicker(config.SoftTime)
	makeRuneTicker = time.NewTicker(config.RuneTime)

	forceStart()

	go togglePauseRune()

	for {
		select {
		case <-foodTicker.C:
			util.RandomSleep(5, time.Second)
			eatFood()

		case <-antiLogoutTicker.C:
			util.RandomSleep(5, time.Second)
			stayLoggedIn()

		case <-ringTicker.C:
			util.RandomSleep(10, time.Second)
			equipRing()

		case <-softTicker.C:
			util.RandomSleep(10, time.Second)
			equipSoft()

		case <-makeRuneTicker.C:
			util.RandomSleep(10, time.Second)
			makeRune()
		}
	}
}

func makeRune() {
	changeWindow()
	logger.Println("Making rune")
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.CreateRune)
		runeCount += 1
		logger.Printf("%d Rune made\n", runeCount)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
	}
}

func equipSoft() {
	changeWindow()
	logger.Println("Equipping Soft")
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipSoft)
		util.RandomSleep(2, time.Second)
	}
}

func equipRing() {
	changeWindow()
	logger.Println("Equipping Ring")
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipRing)
		util.RandomSleep(2, time.Second)
	}
}

func stayLoggedIn() {
	changeWindow()
	logger.Println("Staying logged in")
	robotgo.KeyTap("up")
	util.RandomSleep(5, time.Second)
	robotgo.KeyTap("down")
	util.RandomSleep(3, time.Second)
	robotgo.KeyTap("right")
	util.RandomSleep(6, time.Second)
	robotgo.KeyTap("left")
}

func eatFood() {
	changeWindow()
	logger.Println("Eating food")
	for i := 0; i < 4; i++ {
		robotgo.KeyTap(config.EatFood)
		util.RandomSleep(2, time.Second)
	}
}

func forceStart() {
	equipRing()
	equipSoft()
	eatFood()
	makeRune()
}

func togglePauseRune() {
	events := hook.Start()
	for ev := range events {
		toggled := togglePause(ev)
		if toggled && paused {
			foodTicker.Stop()
			antiLogoutTicker.Stop()
			ringTicker.Stop()
			softTicker.Stop()
			makeRuneTicker.Stop()
			continue
		}

		if toggled && !paused {
			foodTicker.Reset(config.FoodTime)
			antiLogoutTicker.Reset(config.AntiLogoutTime)
			ringTicker.Reset(config.RingTime)
			softTicker.Reset(config.SoftTime)
			makeRuneTicker.Reset(config.RuneTime)
			forceStart()
		}
	}
}
