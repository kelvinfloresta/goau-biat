package util

import (
	"os"
	"os/exec"
	"runtime"
)

var ClearTerminal func()

func init() {
	switch runtime.GOOS {
	case "windows":
		ClearTerminal = func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	case "linux":
		ClearTerminal = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
}
