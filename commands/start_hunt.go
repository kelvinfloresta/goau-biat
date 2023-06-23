package commands

import (
	"goau-biat/config"

	hook "github.com/robotn/gohook"
)

func StartHunt() {
	events := hook.Start()
	for ev := range events {
		togglePause(ev)

		if paused {
			continue
		}

		ScheduleTimer(ev)
		HelpLoot(ev)
		go UltimateHealing(ev)
		SmartRune(ev)
		SafeMagicShield(ev)
		SafeMaxVita(ev)
		Safe("SSA", ev, config.Ssa, config.ExuraVitaKey, config.PotionKey)
		GetLoot(ev)
	}
}
