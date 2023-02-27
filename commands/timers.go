package commands

import (
	"goau-biat/config"
	"goau-biat/sounds"
	"goau-biat/util"
	"log"

	hook "github.com/robotn/gohook"
)

func ScheduleTimer(ev hook.Event) {

	if ev.Keychar == config.UltiamteKey || rune(ev.Rawcode) == config.UltiamteKey {
		log.Default().Printf("Scheduled Ultimate sound")

		util.Scheduler(func() {
			err := sounds.PlaySound(sounds.UltimateReloaded)
			if err != nil {
				log.Default().Printf("Failed to play Ultimate sound")
			}
		}, config.UltimateTime)
	}
}
