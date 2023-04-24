package main

import (
	"fmt"
	"goau-biat/commands"
	"goau-biat/config"
	"goau-biat/sounds"
	"log"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	hook "github.com/robotn/gohook"
	"github.com/tatsushid/go-fastping"
)

func main() {
	var prompt = promptui.Select{
		Label: "Select a command",
		Items: []string{
			"Rune",
			"Loot",
			"Monitor",
		},
	}

	_, command, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	switch command {
	case "Loot":
		log.Default().Println("Starting loot")
		go loot()
	case "Rune":
		log.Default().Println("Starting rune")
		go startRune()
	case "Monitor":
		log.Default().Println("Starting Monitor")
		go startMonitor()
	}

	select {}
}

func startRune() {
	events := hook.Start()
	go commands.Rune()

	for ev := range events {
		commands.Exit(ev)
	}
}

var paused = false

func loot() {
	events := hook.Start()
	for ev := range events {
		togglePause(ev)

		if paused {
			continue
		}

		commands.ScheduleTimer(ev)
		commands.HelpLoot(ev)
		commands.UltimateHealing(ev)
		commands.SmarRune(ev)
		commands.SafeMagicShield(ev)

		if ev.Keychar == config.GetLoot {
			commands.GetLoot(ev)
		}

		if ev.Keychar == config.Checklist {
			sounds.PlaySound(sounds.CheckListLostSouls)
		}

	}
}

func togglePause(ev hook.Event) {
	if ev.Keychar == config.PauseGetLoot || rune(ev.Rawcode) == config.PauseGetLoot {
		log.Default().Println("Pause Get Loot")
		paused = !paused
	}
}

func startMonitor() {
	p := fastping.NewPinger()

	err := p.AddIP("192.168.0.1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	info := color.New(color.FgBlue).SetWriter(log.Default().Writer())
	warning := color.New(color.FgYellow).SetWriter(log.Default().Writer())
	danger := color.New(color.FgRed).SetWriter(log.Default().Writer())
	superDanger := color.New(color.BgHiRed).SetWriter(log.Default().Writer())

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		latency := rtt.Milliseconds()
		switch {
		case latency < 10:
			info.Printf("Good: %s \n", rtt)

		case latency > 20 && latency < 50:
			warning.Printf("Warning: %s \n", rtt)

		case latency > 50 && latency < 100:
			danger.Printf("Danger: %s \n", rtt)

		case latency > 100:
			superDanger.Printf("Danger: %s \n", rtt)
		}
	}

	p.RunLoop()
}
