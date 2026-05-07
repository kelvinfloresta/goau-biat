package config

import (
	"time"
)

const (
	SsaKey                = 'g'
	PotionKey             = "num9"
	GetLoot               = '0'
	Checklist      rune   = '*'
	Pause          rune   = 65481 // ESC 27
	PauseUh        uint16 = 65481 // F12
	CleanHelpLoot  rune   = 65480 // F11
	SetPosHelpLoop        = 'y'
	StopHelpLoot          = 'u'

	Ssa       = 53 // num5
	MightRing = "num8"

	Lang = "pt-br"

	ExuraVitaKey          = "r"
	SecondaryExuraVitaKey = "n"
	SafeMagicShieldKey    = 'f'
	SafeExuraMaxVitaKey   = 'v'

	EatFood        = "f7"
	FoodTime       = 4 * time.Minute
	EquipRing      = "f8"
	RingTime       = 1*time.Minute + 30*time.Second
	EquipSoft      = "f9"
	SoftTime       = 10 * time.Minute
	CreateRune     = "f10"
	RuneTime       = 2 * (1*time.Minute + 15*time.Second)
	AntiLogoutTime = 10 * time.Minute

	X    = 764 + 1530
	Y    = 380
	Diff = 50

	Ctrl_Rawcode uint16 = 162
	A_Rawcode    uint16 = 65
	W_Rawcode    uint16 = 87
	S_Rawcode    uint16 = 83
	D_Rawcode    uint16 = 68
)
