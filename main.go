package main

import (
	"goau-biat/commands"
	"goau-biat/util"

	"github.com/manifoldco/promptui"
)

func main() {
	var prompt = promptui.Select{
		Label: "Select a command",
		Items: []string{
			"Rune",
			"Hunt",
			"Rune with restore window",
		},
	}

	_, command, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	switch command {
	case "Hunt":
		util.Logger.Println("Starting Hunt")
		go commands.StartHunt()
	case "Rune":
		util.Logger.Println("Starting Rune")
		runeMaker := commands.RuneMaker{}
		go runeMaker.Start()

	case "Rune with restore window":
		util.Logger.Println("Starting Rune with restore window")
		runeMaker := commands.RuneMaker{}
		runeMaker.ShouldRestoreWindow = true
		go runeMaker.Start()
	}

	select {}
}
