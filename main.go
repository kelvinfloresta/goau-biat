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
	ev := hook.Start()
	for el := range ev {
		if el.Keychar == config.GetLoot {
			commands.GetLoot(el)
		}

		if el.Keychar == config.PauseGetLoot {
			commands.TogglePauseGetLoot()
		}
	}
}
