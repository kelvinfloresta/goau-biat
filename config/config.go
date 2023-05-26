package config

import (
	"time"
)

const (
	GetLoot               = '-'
	Checklist      rune   = '*'
	Pause          rune   = 27    // ESC
	PauseUh        uint16 = 65481 // F12
	CleanHelpLoot  rune   = 65480 // F11
	SetPosHelpLoop        = 'y'
	StopHelpLoot          = 'u'

	Ssa       = 53 // num5
	MightRing = "num8"

	Lang            = "pt-br"
	UltiamteKey     = '+'
	PreUltimateTime = 10 * time.Second
	UltimateTime    = 40 * time.Second

	ExuraVitaKey       = "v"
	SafeMagicShieldKey = 'f'
	Uh                 = 49 // 1
	UhX                = 1642
	UhY                = 88
	UhYDiff            = 20
	UhRp               = 50 // 2

	AttackRuneKey = 65457

	EatFood        = "f7"
	FoodTime       = 5 * time.Minute
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

	GatewayIP  = "192.168.0.1"
	ExternalIP = "8.8.8.8"
)
