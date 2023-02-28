package commands

import (
	"goau-biat/config"
	"goau-biat/sounds"
	"goau-biat/util"
	"log"
	"time"

	hook "github.com/robotn/gohook"
)

var scheduledUltimate = false

func ScheduleTimer(ev hook.Event) {

	isUltimate := ev.Keychar == config.UltiamteKey || rune(ev.Rawcode) == config.UltiamteKey
	if !scheduledUltimate && isUltimate {
		scheduledUltimate = true
		log.Default().Printf("Scheduled Ultimate sound")

		util.Scheduler(func() {
			err := sounds.PlaySound(sounds.PreUltimateReloaded)
			if err != nil {
				log.Default().Printf("Failed to play Pre Ultimate sound")
			}

			time.Sleep(config.PreUltimateTime)
			err = sounds.PlaySound(sounds.UltimateReloaded)
			if err != nil {
				log.Default().Printf("Failed to play Ultimate sound")
			}

			scheduledUltimate = false
		}, config.UltimateTime-config.PreUltimateTime)
	}
}
