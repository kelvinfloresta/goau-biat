package commands

import (
	"goau-biat/config"
	"goau-biat/util"

	hook "github.com/robotn/gohook"
)

func StartHunt() {
	events := hook.Start()
	for ev := range events {
		if !util.IsClient() {
			continue
		}

		togglePause(ev)

		if paused {
			continue
		}

		SafeMagicShield(ev)
		SafeMaxVita(ev)
		Safe("SSA", ev, config.Ssa, config.ExuraVitaKey, config.PotionKey)
	}
}
