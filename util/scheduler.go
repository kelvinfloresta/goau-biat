package util

import (
	"time"
)

func Scheduler(fn func(), d time.Duration) {
	go (func() {
		time.Sleep(d)
		fn()
	})()
}
