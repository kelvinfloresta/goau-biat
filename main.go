package main

import (
	"fmt"
	"goau-biat/commands"
	"goau-biat/config"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/tatsushid/go-fastping"
)

func main() {
	var prompt = promptui.Select{
		Label: "Select a command",
		Items: []string{
			"Rune",
			"Hunt",
			"Monitor",
		},
	}

	_, command, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	switch command {
	case "Hunt":
		log.Default().Println("Starting Hunt")
		go commands.StartHunt()
	case "Rune":
		log.Default().Println("Starting Rune")
		go commands.StartRune()
	case "Monitor":
		log.Default().Println("Starting Monitor")
		go startMonitor()
	}

	select {}
}

func startMonitor() {
	gatewayPinger := newPinger(config.GatewayIP)
	externalPinger := newPinger(config.ExternalIP)

	var gatewayLatancy time.Duration
	var externalLatency time.Duration

	var printer = newLatencyPrinter()

	gatewayPinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		gatewayLatancy = rtt
		fmt.Println()
		fmt.Print(" Gateway ")
		printer(gatewayLatancy)
		fmt.Print("\n External ")
		printer(externalLatency)
		fmt.Print("\n\n")
	}

	externalPinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		externalLatency = rtt
		fmt.Print(" Gateway ")
		printer(gatewayLatancy)
		fmt.Print("\n External ")
		printer(externalLatency)
		fmt.Print("\n\n")
	}

	gatewayPinger.RunLoop()
	externalPinger.RunLoop()
}

func newLatencyPrinter() func(duration time.Duration) {
	info := color.New(color.FgBlue).SetWriter(log.Writer())
	warning := color.New(color.FgYellow).SetWriter(log.Writer())
	danger := color.New(color.FgRed).SetWriter(log.Writer())
	superDanger := color.New(color.BgHiRed).SetWriter(log.Writer())

	return func(duration time.Duration) {
		latency := duration.Milliseconds()

		switch {
		case latency < 10:
			info.Printf("%s", duration)

		case latency > 20 && latency < 50:
			warning.Printf("%s", duration)

		case latency > 50 && latency < 100:
			danger.Printf("%s", duration)

		case latency > 100:
			superDanger.Printf("%s", duration)
		}
	}

}

func newPinger(ip string) *fastping.Pinger {
	p := fastping.NewPinger()

	err := p.AddIP(ip)
	if err != nil {
		log.Default().Fatalln(err)
	}
	return p
}
