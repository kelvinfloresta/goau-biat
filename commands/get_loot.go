package commands

import (
	"fmt"
	"goau-biat/config"
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var paused = false
var x = 820
var y = 520
var diff = 70

func GetLoot(el hook.Event) {
	if paused {
		return
	}

	robotgo.KeyDown("alt")
	moveClick(x-diff, y)
	moveClick(x-diff, y-diff)
	moveClick(x, y-diff)
	moveClick(x+diff, y-diff)
	moveClick(x+diff, y)
	moveClick(x+diff, y+diff)
	moveClick(x, y+diff)
	moveClick(x-diff, y+diff)
	robotgo.KeyUp("alt")
}

func HelpLoot(el hook.Event) {
	list := []string{}

	if el.Keychar == config.CleanHelpLoot {
		log.Default().Println("Clean Help Loot")
		list = []string{}
	}

	if el.Keychar == config.SetPosHelpLoop || rune(el.Rawcode) == config.SetPosHelpLoop {
		x, y := robotgo.GetMousePos()
		log.Default().Printf("Set Help Loot: %d %d\n", x, y)
		list = append(list, fmt.Sprintf("%d %d", x, y))
	}

	if el.Keychar == config.StopHelpLoot || rune(el.Rawcode) == config.StopHelpLoot {
		log.Default().Println(strings.Join(list, ","))
	}

	if el.Keychar == config.PauseGetLoot || rune(el.Rawcode) == config.PauseGetLoot {
		log.Default().Println("Pause Get Loot")
		paused = !paused
	}
}
