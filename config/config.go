package config

import (
	"time"
)

const (
	GetLoot             = '-'
	Checklist      rune = '*'
	Pause          rune = 27    // ESC
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
	UhX                = 1688
	UhY                = 596
	UhRp               = 52 // 4

	AttackRuneKey = 49

	EatFood        = "f7"
	FoodTime       = 10 * time.Minute
	EquipRing      = "f8"
	RingTime       = 1 * time.Minute
	EquipSoft      = "f9"
	SoftTime       = 10 * time.Minute
	CreateRune     = "f10"
	RuneTime       = 5 * time.Minute
	AntiLogoutTime = 10 * time.Minute

	X    = 870
	Y    = 530
	Diff = 70
)
