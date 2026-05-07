package util

import (
	"log"
	"os"
)

var Logger = log.New(os.Stderr, "", 4)
