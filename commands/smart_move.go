package commands

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func SmartMove() {
	bit := robotgo.CaptureScreen(10, 20, 30, 40)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bit)

	fmt.Println("bitmap...", bit)
}
