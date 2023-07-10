package commands

import (
	"fmt"
	"goau-biat/config"
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func GetLoot(ev hook.Event) {
	if ev.Keychar != config.GetLoot {
		return
	}

	robotgo.KeyDown("alt")
	moveClick(config.X-config.Diff, config.Y)
	moveClick(config.X-config.Diff, config.Y-config.Diff)
	moveClick(config.X, config.Y-config.Diff)
	moveClick(config.X+config.Diff, config.Y-config.Diff)
	moveClick(config.X+config.Diff, config.Y)
	moveClick(config.X+config.Diff, config.Y+config.Diff)
	moveClick(config.X, config.Y+config.Diff)
	moveClick(config.X-config.Diff, config.Y+config.Diff)
	robotgo.KeyUp("alt")
}

func HelpLoot(el hook.Event) {
	list := []string{}

	if el.Keychar == config.CleanHelpLoot {
		log.Default().Println("Clean Help Loot")
		list = []string{}
	}

	if el.Keychar == config.SetPosHelpLoop || rune(el.Rawcode) == config.SetPosHelpLoop {
		x, y := robotgo.Location()
		log.Default().Printf("Set Help Loot: %d %d\n", x, y)
		list = append(list, fmt.Sprintf("%d %d", x, y))
	}

	if el.Keychar == config.StopHelpLoot || rune(el.Rawcode) == config.StopHelpLoot {
		log.Default().Println(strings.Join(list, ","))
	}
}
