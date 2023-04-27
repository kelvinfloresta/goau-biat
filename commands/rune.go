package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func Rune() {
	food := time.NewTicker(config.FoodTime)
	antiLogout := time.NewTicker(config.AntiLogoutTime)
	ring := time.NewTicker(config.RingTime)
	soft := time.NewTicker(config.SoftTime)
	runes := time.NewTicker(config.RuneTime)

	equipRing()
	equipSoft()
	eatFood()
	makeRune()

	for {
		select {
		case <-food.C:
			util.RandomSleep(5, time.Second)
			eatFood()

		case <-antiLogout.C:
			util.RandomSleep(5, time.Second)
			stayLoggedIn()

		case <-ring.C:
			util.RandomSleep(10, time.Second)
			equipRing()

		case <-soft.C:
			util.RandomSleep(10, time.Second)
			equipSoft()

		case <-runes.C:
			util.RandomSleep(10, time.Second)
			makeRune()
		}
	}
}

func makeRune() {
	changeWindow()
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.CreateRune)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
	}
}

func equipSoft() {
	changeWindow()
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipSoft)
		util.RandomSleep(2, time.Second)
	}
}

func equipRing() {
	changeWindow()
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipRing)
		util.RandomSleep(2, time.Second)
	}
}

func stayLoggedIn() {
	changeWindow()
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
	for i := 0; i < 4; i++ {
		robotgo.KeyTap(config.EatFood)
		util.RandomSleep(2, time.Second)
	}
}
