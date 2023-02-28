package commands

import (
	"goau-biat/config"
	"goau-biat/util"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func changeWindow() {
	robotgo.ActivePID(21366)
}

func Rune() {
	time.Sleep(5 * time.Second)

	util.PermaScheduler(func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		changeWindow()
		robotgo.KeyTap(config.EatFood)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		robotgo.KeyTap(config.EatFood)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		robotgo.KeyTap(config.EatFood)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		robotgo.KeyTap(config.EatFood)
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
		robotgo.KeyTap(config.CreateRune)
	}, config.RuneTime)

}
