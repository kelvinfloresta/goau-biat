package util

import (
	"math/rand"
	"time"
)

func RandomSleep(randomValue int, duration time.Duration) {
	time.Sleep(1 + time.Duration(rand.Intn(randomValue-1))*duration)
}
