package config

import "time"

const (
	GetLoot             = '-'
	PauseGetLoot   rune = 65481 // F12
	CleanHelpLoot  rune = 65480 // F11
	SetPosHelpLoop      = 'y'
	StopHelpLoot        = 'u'

	Lang            = "pt-br"
	UltiamteKey     = '+'
	PreUltimateTime = 10 * time.Second
	UltimateTime    = 5 * time.Second

	EatFood        = "f7"
	FoodTime       = 10 * time.Minute
	EquipRing      = "f8"
	RingTime       = 5 * time.Minute
	EquipSoft      = "f9"
	SoftTime       = 1 * time.Hour
	CreateRune     = "f10"
	RuneTime       = 1 * time.Minute
	AntiLogoutTime = 10 * time.Minute

	X    = 870
	Y    = 530
	Diff = 70
)
