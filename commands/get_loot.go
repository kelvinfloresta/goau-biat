package commands

import (
	"fmt"
	"goau-biat/config"
	"strings"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var paused = false
var x = 590
var y = 375

func GetLoot(el hook.Event) {
	if paused {
		return
	}

	robotgo.KeyDown("alt")
	moveClick(x-50, y)
	moveClick(x-50, y-50)
	moveClick(x, y-50)
	moveClick(x+50, y-50)
	moveClick(x+50, y)
	moveClick(x+50, y+50)
	moveClick(x, y+50)
	moveClick(x-50, y+50)
	robotgo.KeyUp("alt")
}

func TogglePauseGetLoot() {
	paused = !paused
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
}
