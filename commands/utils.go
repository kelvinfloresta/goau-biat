package commands

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

func moveClick(x, y int, args ...any) {
	randX := rand.Intn(10) + x
	randY := rand.Intn(10) + y

	robotgo.MoveClick(randX, randY, args...)
}
