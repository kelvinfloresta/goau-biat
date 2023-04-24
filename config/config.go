package config

import (
	"time"
)

const (
	GetLoot             = '-'
	Checklist      rune = '*'
	PauseGetLoot   rune = 65481 // F12
	CleanHelpLoot  rune = 65480 // F11
	SetPosHelpLoop      = 'y'
	StopHelpLoot        = 'u'

	Ssa       = 53 // num5
	MightRing = "num8"

	Lang            = "pt-br"
	UltiamteKey     = '+'
	PreUltimateTime = 10 * time.Second
	UltimateTime    = 40 * time.Second

	ExuraVitaKey       = "v"
	SafeMagicShieldKey = 'f'
	Uh                 = 50 // 2
	AttackRuneKey = 49

	CancelRunes    = 27 // ESC
	EatFood        = "f7"
	FoodTime       = 10 * time.Minute
	EquipRing      = "f8"
	RingTime       = 1 * time.Minute
	EquipSoft      = "f9"
	SoftTime       = 1 * time.Hour
	CreateRune     = "f10"
	RuneTime       = 5 * time.Minute
	AntiLogoutTime = 10 * time.Minute

	X    = 870
	Y    = 530
	Diff = 70
)
