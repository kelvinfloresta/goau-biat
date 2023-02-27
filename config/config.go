package config

import "time"

const (
	GetLoot             = '-'
	PauseGetLoot   rune = 65481 // F12
	CleanHelpLoot  rune = 65480 // F11
	SetPosHelpLoop      = 'y'
	StopHelpLoot        = 'u'

	Lang         = "pt-br"
	UltiamteKey  = '+'
	UltimateTime = 28 * time.Second

	X    = 870
	Y    = 530
	Diff = 70
)
