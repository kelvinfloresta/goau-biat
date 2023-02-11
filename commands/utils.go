package commands

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

func moveClick(x, y int, args ...any) {
	randX := rand.Intn(5) + x
	randY := rand.Intn(5) + y

	robotgo.MoveClick(randX, randY, args...)
}
