package commands

import (
	"goau-biat/config"
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
			eatFood()

		case <-antiLogout.C:
			stayLoggedIn()

		case <-ring.C:
			equipRing()

		case <-soft.C:
			equipSoft()

		case <-runes.C:
			makeRune()
		}
	}
}

func makeRune() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	changeWindow()

	for i := 0; i < 5; i++ {
		robotgo.KeyTap(config.CreateRune)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
	}
}

func equipSoft() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	changeWindow()
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipSoft)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
	}
}

func equipRing() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	changeWindow()
	for i := 0; i < 2; i++ {
		robotgo.KeyTap(config.EquipRing)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
	}
}

func stayLoggedIn() {
	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
	changeWindow()
	robotgo.KeyTap("up")
	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
	robotgo.KeyTap("down")
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	robotgo.KeyTap("right")
	time.Sleep(time.Duration(1+rand.Intn(6)) * time.Second)
	robotgo.KeyTap("left")
}

func eatFood() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	changeWindow()
	for i := 0; i < 4; i++ {
		robotgo.KeyTap(config.EatFood)
		time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
	}
}
