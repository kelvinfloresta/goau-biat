package commands

import (
	"fmt"
	"goau-biat/util"
	"log"
	"strings"
	"time"

	"goau-biat/util/winapi"
	ps "github.com/mitchellh/go-ps"
)

var clientPid = getPid()

func getPid() int {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		panic(err)
	}

	for _, process := range processList {
		if strings.Contains(process.Executable(), "client.exe") &&
			!strings.Contains(process.Executable(), "ts3") &&
			!strings.Contains(process.Executable(), "launcher") {
			fmt.Println(process.Pid())
			return process.Pid()
		}
	}

	return 0
}

func ChangeWindow() {
	util.Logger.Println("Changing to client window")
	winapi.ActivePid(clientPid)
	time.Sleep(1 * time.Second)
}
