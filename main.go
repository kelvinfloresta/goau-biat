package main

import (
	"goau-biat/commands"
	"goau-biat/config"
	"log"

	"github.com/manifoldco/promptui"
	hook "github.com/robotn/gohook"
)

func main() {
	var prompt = promptui.Select{
		Label: "Select a command",
		Items: []string{
			"Rune",
			"Loot",
		},
	}

	_, command, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	switch command {
	case "Loot":
		log.Default().Println("Starting loot")
		loot()
	case "Rune":
		log.Default().Println("Starting rune")
		startRune()
	}
}

func startRune() {
	commands.Rune()
	select {}
}

func loot() {
	events := hook.Start()
	for ev := range events {
		commands.ScheduleTimer(ev)
		commands.HelpLoot(ev)
		if ev.Keychar == config.GetLoot {
			commands.GetLoot(ev)
		}
	}
}
