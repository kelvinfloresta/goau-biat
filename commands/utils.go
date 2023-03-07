package commands

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	ps "github.com/mitchellh/go-ps"
)

func moveClick(x, y int, args ...any) {
	randX := rand.Intn(5) + x
	randY := rand.Intn(5) + y

	robotgo.MoveClick(randX, randY, args...)
}

var pid = getPid()

func getPid() int32 {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		panic(err)
	}

	for _, process := range processList {
		if strings.Contains(process.Executable(), "client") {
			return int32(process.Pid())
		}
	}

	return 0
}

func changeWindow() {
	robotgo.ActivePID(pid)
	time.Sleep(1 * time.Second)
}
