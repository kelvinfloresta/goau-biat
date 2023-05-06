package commands

import (
	"goau-biat/config"
	"goau-biat/sounds"

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
		UltimateHealing(ev)
		SmarRune(ev)
		SafeMagicShield(ev)

		if ev.Keychar == config.GetLoot {
			GetLoot(ev)
		}

		if ev.Keychar == config.Checklist {
			sounds.PlaySound(sounds.CheckListGoannas)
		}

	}
}
