package main

import (
	"goau-biat/commands"
	"goau-biat/config"

	hook "github.com/robotn/gohook"
)

func main() {
	test()
}

func test() {
	events := hook.Start()

	for ev := range events {
		commands.ScheduleTimer(ev)
		commands.HelpLoot(ev)
		if ev.Keychar == config.GetLoot {
			commands.GetLoot(ev)
		}
	}
}
