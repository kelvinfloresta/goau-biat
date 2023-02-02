package commands

import (
	"fmt"
	"goau-biat/config"
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
		list = []string{}
	}

	if el.Keychar == config.SetPosHelpLoop {
		x, y := robotgo.GetMousePos()
		fmt.Printf("Set: %d %d\n", x, y)
		list = append(list, fmt.Sprintf("%d %d", x, y))
	}

	if el.Keychar == config.StopHelpLoot {
		fmt.Println(strings.Join(list, ","))
	}

	if el.Keychar == config.PauseGetLoot {
		paused = !paused
	}
}
