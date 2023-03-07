package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func Rune() {

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		changeWindow()
		for i := 0; i < 5; i++ {
			robotgo.KeyTap(config.EatFood)
			time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
		}
	}, config.FoodTime)

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		changeWindow()
		for i := 0; i < 5; i++ {
			robotgo.KeyTap(config.EatFood)
			time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
		}
	}, config.FoodTime)

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
		changeWindow()
		robotgo.KeyTap("up")
		time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
		robotgo.KeyTap("down")
		time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
		robotgo.KeyTap("right")
		time.Sleep(time.Duration(1+rand.Intn(6)) * time.Second)
		robotgo.KeyTap("left")
	}, config.AntiLogoutTime)

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		changeWindow()
		robotgo.KeyTap(config.EquipRing)
	}, config.RingTime)

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		changeWindow()
		robotgo.KeyTap(config.EquipSoft)
	}, config.SoftTime)

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		changeWindow()

		for i := 0; i < 5; i++ {
			robotgo.KeyTap(config.CreateRune)
			time.Sleep((2 + time.Duration(rand.Intn(2))) * time.Second)
		}

	}, config.RuneTime)

}
